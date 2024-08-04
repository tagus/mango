package mango

import (
	"log"
	"log/slog"
	"os"
	"runtime/debug"
)

func Init(level slog.Level, application string) {
	var h slog.Handler
	h = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	h = h.WithAttrs([]slog.Attr{
		slog.String("app", application),
	})
	logger := slog.New(h)
	slog.SetDefault(logger)
}

func Fatal(v ...any) {
	debug.PrintStack()
	log.Fatal(v...)
}

func FatalIf(err error, v ...any) {
	if err == nil {
		return
	}
	args := append([]any{err}, v...)
	Fatal(args...)
}

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
