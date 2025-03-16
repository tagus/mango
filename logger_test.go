package mango

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log/slog"
	"os"
	"sync"
	"testing"
	"time"
)

func TestWrappingStackError(t *testing.T) {
	testErr := errors.New("test error")
	err := WrapError(testErr)

	testErr = fmt.Errorf("wrapped error: %w", err)

	var stErr *StackError
	require.Equal(t, true, errors.As(testErr, &stErr))
}

func TestLogHandler_Enabled(t *testing.T) {
	ctx := context.TODO()

	tests := []struct {
		name     string
		curLevel slog.Level
		level    slog.Level
		expected bool
	}{
		{
			name:     "debug < info",
			curLevel: slog.LevelDebug,
			level:    slog.LevelInfo,
			expected: true,
		},
		{
			name:     "info < error",
			curLevel: slog.LevelError,
			level:    slog.LevelInfo,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lh := &LogHandler{
				app:   "test-app",
				level: test.curLevel,
				out:   os.Stdout,
				mu:    &sync.Mutex{},
			}

			actual := lh.Enabled(ctx, test.level)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestLogHandler(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	lh := &LogHandler{
		app:   "test-app",
		level: slog.LevelDebug,
		out:   buf,
		mu:    &sync.Mutex{},
	}
	logger := slog.New(lh).With("foo", "bar")

	ts := time.Date(2025, 1, 19, 21, 49, 1, 0, time.UTC)
	logger.Warn("warn message", "ts", ts, "count", 2)

	output := buf.String()
	expected := `[[38;2;241;196;15mWARN[0m][test-app] warn message
  foo: bar
  ts: 2025-01-19 21:49:01
  count: 2
`
	assert.Equal(t, expected, output)
}
