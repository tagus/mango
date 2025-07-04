package mango

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatSimpleDate(t *testing.T) {
	ts := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "01-01-21", FormatSimpleDate(&ts))
}

func TestFormatTimeSince(t *testing.T) {
	ts := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	timeSince := FormatTimeSince(&ts)
	assert.NotEmpty(t, timeSince)
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  time.Time
		expectErr bool
	}{
		{
			name:     "null value",
			input:    `null`,
			expected: time.Time{},
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: time.Time{},
		},
		{
			name:     "valid RFC3339",
			input:    `"2023-10-01T12:00:00Z"`,
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "valid date only",
			input:    `"2023-10-01"`,
			expected: time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "valid date time",
			input:    `"2023-10-01 12:00:00"`,
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.Local),
		},
		{
			name:      "invalid value",
			input:     `"invalid-date"`,
			expectErr: true,
		},
		{
			name:      "invalid format",
			input:     `"Jan 01 2006"`,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ts Timestamp
			err := ts.UnmarshalJSON([]byte(test.input))
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expected, ts.Time)
		})
	}
}

func TestTimestamp_Scan(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  time.Time
		expectErr bool
	}{
		{
			name:     "nil value",
			input:    nil,
			expected: time.Time{},
		},
		{
			name:     "time.Time value",
			input:    time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "valid string value",
			input:    "2023-10-01T12:00:00Z",
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name:      "invalid type",
			input:     12345,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ts Timestamp
			err := ts.Scan(test.input)
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expected, ts.Time)
		})
	}
}

func TestTimestamp_Value(t *testing.T) {
	tests := []struct {
		name     string
		input    Timestamp
		expected any
	}{
		{
			name:     "zero value",
			input:    Timestamp{Time: time.Time{}},
			expected: nil,
		},
		{
			name:     "non-zero value",
			input:    Timestamp{Time: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)},
			expected: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := test.input.Value()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:  "valid RFC3339",
			input: "2023-10-01T12:00:00Z",
		},
		{
			name:  "valid date only",
			input: "2023-10-01",
		},
		{
			name:  "valid date time",
			input: "2023-10-01 12:00:00",
		},
		{
			name:      "invalid format",
			input:     "invalid-date",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts, err := ParseTimestamp(test.input)
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotZero(t, ts.Time)
		})
	}
}
