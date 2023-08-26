package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	lst := []int{1, 2, 3}

	expected := []int{2, 4, 6}
	actual := Map(lst, func(i int) int { return i * 2 })

	assert.Equal(t, expected, actual)
}

func TestFilter(t *testing.T) {
	lst := []int{1, 2, 3}

	expected := []int{2}
	actual := Filter(lst, func(i int) bool { return i%2 == 0 })

	assert.Equal(t, expected, actual)
}
