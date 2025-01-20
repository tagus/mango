package mango

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"
	"sync"
)

// Init creates a slog logger instance with a custom LogHandler
// and sets it as the default logger
func Init(level slog.Level, application string) {
	var h slog.Handler
	h = &LogHandler{
		level: level,
		app:   application,
		out:   os.Stdout,
		mu:    &sync.Mutex{},
	}
	logger := slog.New(h)
	slog.SetDefault(logger)
}

// LogHandler is a custom slog handler that outputs to stdout
type LogHandler struct {
	app   string
	level slog.Leveler
	out   io.Writer
	attrs []slog.Attr

	// NOTE: since the WithAttrs and WithGroup makes copies of the LogHandler,
	// we need to ensure that the same mutex instance is used across all instances
	mu *sync.Mutex
}

// Enabled only records the log if the record's log level is greater than
// or equal to the handlers log level
func (h *LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

// Handle outputs the log record to stdout. Since this logger is meant to be used for
// local applications with a small amount of output, it will not log the timestamps.
func (h *LogHandler) Handle(ctx context.Context, record slog.Record) error {
	var builder strings.Builder

	_, err := builder.WriteString(fmt.Sprintf("[%s][%s] %s", h.getLevelPrefix(record.Level), h.app, record.Message))
	if err != nil {
		return err
	}

	hasAttrs := len(h.attrs) > 0 || record.NumAttrs() > 0
	if hasAttrs {
		_, err := builder.WriteString("\n")
		if err != nil {
			return err
		}
		for _, attr := range h.attrs {
			if err := h.appendAttr(&builder, attr); err != nil {
				return err
			}
		}
		var cbError error
		record.Attrs(func(attr slog.Attr) bool {
			if err := h.appendAttr(&builder, attr); err != nil {
				cbError = err
				return false
			}
			return true
		})
		if cbError != nil {
			return cbError
		}
	}

	if !hasAttrs {
		_, err = builder.WriteString("\n")
		if err != nil {
			return err
		}
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err = h.out.Write([]byte(builder.String()))
	return err
}

func (h *LogHandler) appendAttr(builder *strings.Builder, attr slog.Attr) error {
	// ignore an empty attribute
	if attr.Equal(slog.Attr{}) {
		return nil
	}
	var value string
	switch attr.Value.Kind() {
	case slog.KindString:
		value = attr.Value.String()
	case slog.KindTime:
		value = attr.Value.Time().Format("2006-01-02 15:04:05")
	default:
		value = attr.Value.String()
	}
	builder.WriteString(fmt.Sprintf("  %s: %s\n", attr.Key, value))
	return nil
}

func (h *LogHandler) getLevelPrefix(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return fmt.Sprintf("%sDEBUG%s", Blue, Reset)
	case slog.LevelInfo:
		return fmt.Sprintf("%sINFO%s", Green, Reset)
	case slog.LevelWarn:
		return fmt.Sprintf("%sWARN%s", Yellow, Reset)
	case slog.LevelError:
		return fmt.Sprintf("%sERROR%s", Red, Reset)
	default:
		return fmt.Sprintf("%sUNKNOWN%s", Purple, Reset)
	}
}

func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attrs = append(h.attrs, attrs...)
	return h
}

// NOTE: groups are not currently supported
func (h *LogHandler) WithGroup(name string) slog.Handler {
	return h
}

/******************************************************************************/

func Fatal(err error, args ...any) {
	debug.PrintStack()
	slog.Error(err.Error(), args...)
	os.Exit(1)
}

func FatalIf(err error, v ...any) {
	if err == nil {
		return
	}
	Fatal(err, v...)
}

/******************************************************************************/

// StackError wraps an error with the stack information
// where the error occurred. Note that since stack is saved during creation
// we need to create this at the exact location where the base error is thrown
// and not at a higher level.
type StackError struct {
	base  error
	stack []byte
}

func (s *StackError) Error() string {
	return s.base.Error()
}

func (s *StackError) Unwrap() error {
	return s.base
}

func WrapError(err error) error {
	st := &StackError{
		base:  err,
		stack: debug.Stack(),
	}
	return st
}
