package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int
		expectError bool
	}{
		{
			name:        "positive number",
			input:       "42",
			expected:    42,
			expectError: false,
		},
		{
			name:        "negative number",
			input:       "-7",
			expected:    -7,
			expectError: false,
		},
		{
			name:        "zero",
			input:       "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "not a number",
			input:       "notanumber",
			expected:    0,
			expectError: true,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    0,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := ParseInt(test.input)
			if test.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		})
	}
}
func TestParseInt64(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int64
		expectError bool
	}{
		{
			name:        "max int64",
			input:       "9223372036854775807",
			expected:    9223372036854775807,
			expectError: false,
		},
		{
			name:        "min int64",
			input:       "-9223372036854775808",
			expected:    -9223372036854775808,
			expectError: false,
		},
		{
			name:        "zero",
			input:       "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "large positive number",
			input:       "1234567890123456789",
			expected:    1234567890123456789,
			expectError: false,
		},
		{
			name:        "not a number",
			input:       "notanumber",
			expected:    0,
			expectError: true,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    0,
			expectError: true,
		},
		{
			name:        "overflow int64",
			input:       "18446744073709551615",
			expected:    0,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := ParseInt64(test.input)
			if test.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		})
	}
}
