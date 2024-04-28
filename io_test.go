package mango

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ReadAllString(t *testing.T) {
	rdr := bytes.NewBufferString("hello world")
	val, err := ReadAllString(rdr)
	require.NoError(t, err)
	require.Equal(t, "hello world", val)
}
