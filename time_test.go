package mango

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
