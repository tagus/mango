package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvStringSlice(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		defaultValue []string
		expected     []string
	}{
		{
			name:         "empty env variable",
			value:        "",
			defaultValue: []string{"default1", "default2"},
			expected:     []string{"default1", "default2"},
		},
		{
			name:         "env variable with values",
			value:        "value1,value2,value3",
			defaultValue: []string{"default1", "default2"},
			expected:     []string{"value1", "value2", "value3"},
		},
		{
			name:         "env variable with spaces",
			value:        " value1 , value2 , value3 ",
			defaultValue: []string{"default1", "default2"},
			expected:     []string{"value1", "value2", "value3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TEST_ENV", test.value)
			result := GetEnvStringSlice("TEST_ENV", test.defaultValue)
			assert.Equal(t, test.expected, result)
		})
	}
}
