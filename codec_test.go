package mango

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type CodecTest struct {
	Foo string `json:"foo"`
}

func TestMarshal(t *testing.T) {
	ct := CodecTest{Foo: "bar"}

	rdr, err := Marshal(ct)
	require.NoError(t, err)

	buf, err := io.ReadAll(rdr)
	require.NoError(t, err)
	require.Equal(t, "{\"foo\":\"bar\"}\n", string(buf))
}

func TestMarshalToString(t *testing.T) {
	ct := CodecTest{Foo: "bar"}

	str, err := MarshalToString(ct)
	require.NoError(t, err)

	require.Equal(t, "{\"foo\":\"bar\"}\n", str)
}

func TestUnmarshal(t *testing.T) {
	rdr := strings.NewReader(`{"foo":"bar"}`)
	decoded, err := Unmarshal[CodecTest](rdr)

	require.NoError(t, err)
	require.NotNil(t, decoded)
	require.Equal(t, CodecTest{Foo: "bar"}, *decoded)
}

func TestUnmarshalFromString(t *testing.T) {
	decoded, err := UnmarshalFromString[CodecTest](`{"foo":"bar"}`)
	require.NoError(t, err)
	require.NotNil(t, decoded)
	require.Equal(t, CodecTest{Foo: "bar"}, *decoded)

	decodedSlice, err := UnmarshalFromString[[]string](`["foo","bar"]`)
	require.NoError(t, err)
	require.NotNil(t, decodedSlice)
	require.Equal(t, []string{"foo", "bar"}, *decodedSlice)
}
