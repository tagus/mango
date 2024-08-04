package mango

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWrappingStackError(t *testing.T) {
	testErr := errors.New("test error")
	err := WrapError(testErr)

	testErr = fmt.Errorf("wrapped error: %w", err)

	var stErr *StackError
	require.Equal(t, true, errors.As(testErr, &stErr))
}
