package mango

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CodecTest struct {
	Foo string `json:"foo"`
}

func TestMarshal(t *testing.T) {
	ct := CodecTest{Foo: "bar"}

	rdr, err := Marshal(ct)
	assert.NoError(t, err)

	buf, err := io.ReadAll(rdr)
	assert.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\"}\n", string(buf))
}

func TestUnmarshal(t *testing.T) {
	rdr := io.NopCloser(strings.NewReader(`{"foo":"bar"}`))
	decoded, err := Unmarshal[CodecTest](rdr)

	assert.NoError(t, err)
	assert.NotNil(t, decoded)
	assert.Equal(t, CodecTest{Foo: "bar"}, *decoded)
}
