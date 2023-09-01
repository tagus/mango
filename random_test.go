package mango

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortID(t *testing.T) {
	id := ShortID()
	assert.NotEmpty(t, id)

	re := regexp.MustCompile("[a-z0-9]+")
	assert.True(t, re.Match([]byte(id)))
}
