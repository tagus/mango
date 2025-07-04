package mango

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_withoutUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	ctx := context.TODO()
	_, err := NewRequest().
		SetClient(server.Client()).
		Do(ctx)

	assert.Error(t, err)
}

func TestRequest_withoutMethod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	ctx := context.TODO()
	_, err := NewRequest().
		SetClient(server.Client()).
		SetUrl(server.URL).
		Do(ctx)

	assert.Error(t, err)
}

func TestRequest_post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		buf, _ := io.ReadAll(req.Body)
		assert.Equal(t, "{\"foo\":\"bar\"}\n", string(buf))
		assert.Equal(t, "POST", req.Method)
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	ctx := context.TODO()
	res, err := NewRequest().
		SetMethod("POST").
		SetClient(server.Client()).
		SetUrl(server.URL).
		AddRequestField("foo", "bar").
		Do(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, res.Body)
}

func TestRequest_withParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "bar", req.URL.Query().Get("foo"))
		w.Write([]byte("ok"))
	}))
	defer server.Close()

	ctx := context.TODO()
	res, err := NewRequest().
		Post().
		SetClient(server.Client()).
		SetUrl(server.URL).
		AddQueryParam("foo", "bar").
		Do(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, res.Body)
}
