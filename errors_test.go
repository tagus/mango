package mango

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidationError_Is(t *testing.T) {
	tests := []struct {
		name     string
		target   error
		expected bool
	}{
		{
			name:     "matches validation sentinel",
			target:   ErrValidationErr,
			expected: true,
		},
		{
			name:     "does not match other error",
			target:   errors.New("other"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := errors.Is(ValidationError("bad input"), tt.target)
			require.Equal(t, tt.expected, actual)
		})
	}
}
