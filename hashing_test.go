package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashStrings(t *testing.T) {
	res, err := HashStrings("foo", "bar")
	assert.NoError(t, err)
	assert.Equal(t, "3214735720", res)
}
