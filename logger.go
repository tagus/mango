package mango

import (
	"context"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
)

func Init(level slog.Level, prefix string) {
	h := &LogHandler{
		prefix: prefix,
		handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	}
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

/******************************************************************************/

type LogHandler struct {
	handler slog.Handler
	prefix  string
}

func (h *LogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.handler.WithAttrs(attrs)
}

func (h *LogHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}
