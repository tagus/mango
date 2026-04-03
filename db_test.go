package mango

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOrderBy(t *testing.T) {
	tests := []struct {
		name          string
		orderBy       string
		allowedFields []string
		expected      []string
		expectedErr   bool
	}{
		{
			name:          "mixed directions and default desc",
			orderBy:       "foo desc,bar,baz asc",
			allowedFields: []string{"foo", "bar", "baz"},
			expected:      []string{"foo desc", "bar desc", "baz asc"},
		},
		{
			name:          "invalid field",
			orderBy:       "foo desc,qux",
			allowedFields: []string{"foo", "bar"},
			expectedErr:   true,
		},
		{
			name:          "invalid direction",
			orderBy:       "foo up",
			allowedFields: []string{"foo"},
			expectedErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := ParseOrderBy(tt.orderBy, tt.allowedFields)
			if tt.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}
