package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFind(t *testing.T) {
	lst := []int{1, 2, 3}

	actual, ok := Find(lst, func(i int) bool { return i == 2 })
	require.True(t, ok)
	require.Equal(t, 2, actual)

	actual, ok = Find(lst, func(i int) bool { return i == 45 })
	require.False(t, ok)
	require.Equal(t, 0, actual)
}

func TestSliceEqual(t *testing.T) {
	first := []int{1, 2, 3}
	second := []int{1, 2, 3}

	assert.True(t, SliceEqual(first, second, func(a, b int) bool { return a == b }))
	assert.False(t, SliceEqual(first, []int{1, 2}, func(a, b int) bool { return a == b }))
}

func TestStringSliceEqual(t *testing.T) {
	first := []string{"hello", "world"}
	second := []string{"hello", "world"}

	assert.True(t, StringSliceEqual(first, second))
	assert.False(t, StringSliceEqual(first, []string{"hello"}))
}
