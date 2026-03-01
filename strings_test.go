package mango

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			actual, err := ParseInt(test.input)
			if test.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
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
			actual, err := ParseInt64(test.input)
			if test.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestParseListInt(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []int
		expectError bool
	}{
		{
			name:        "comma separated ints",
			input:       "1,2,3",
			expected:    []int{1, 2, 3},
			expectError: false,
		},
		{
			name:        "comma separated ints with spaces",
			input:       "1, 2, 3",
			expected:    []int{1, 2, 3},
			expectError: false,
		},
		{
			name:        "empty input",
			input:       "",
			expected:    []int{},
			expectError: false,
		},
		{
			name:        "filters empty entries",
			input:       "1, ,2,,3",
			expected:    []int{1, 2, 3},
			expectError: false,
		},
		{
			name:        "invalid int in list",
			input:       "1,abc,3",
			expected:    nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := ParseList(test.input, ParseInt)
			if test.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestParseListString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "comma separated strings",
			input:    "a,b,c",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "comma separated strings with spaces",
			input:    "a, b, c",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := ParseList(test.input, ParseString)
			require.NoError(t, err)
			require.Equal(t, test.expected, actual)
		})
	}
}
