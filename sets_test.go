package mango

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Has(t *testing.T) {
	st := NewSet(0, 1, 2)

	assert.True(t, st.Has(0))
	assert.True(t, st.Has(0))
	assert.False(t, st.Has(3))
}

func TestSet_Add(t *testing.T) {
	st := NewSet(0, 1, 2)

	assert.False(t, st.Has(3))
	st.Add(3)
	assert.True(t, st.Has(3))
}

func TestSet_Delete(t *testing.T) {
	st := NewSet(0, 1, 2)

	assert.True(t, st.Has(2))
	st.Delete(2)
	assert.False(t, st.Has(2))
}

func TestSet_Size(t *testing.T) {
	st := NewSet(0, 1, 2)

	assert.Equal(t, 3, st.Size())
	st.Delete(2)
	assert.Equal(t, 2, st.Size())
}
