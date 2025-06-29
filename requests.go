package mango

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type RequestBuilder struct {
	url     string
	method  string
	payload map[string]any
	params  map[string]string
	client  *http.Client
}

func NewRequest() *RequestBuilder {
	return &RequestBuilder{
		params:  make(map[string]string),
		payload: make(map[string]any),
	}
}

func (rb *RequestBuilder) SetUrl(url string) *RequestBuilder {
	rb.url = url
	return rb
}

func (rb *RequestBuilder) AddRequestField(key string, value any) *RequestBuilder {
	rb.payload[key] = value
	return rb
}

func (rb *RequestBuilder) AddRequestFieldOptional(key string, value any, predicate bool) *RequestBuilder {
	if !predicate {
		return rb
	}
	rb.payload[key] = value
	return rb
}

func (rb *RequestBuilder) AddQueryParam(key string, value any) *RequestBuilder {
	rb.params[key] = fmt.Sprintf("%v", value)
	return rb
}

func (rb *RequestBuilder) SetClient(client *http.Client) *RequestBuilder {
	rb.client = client
	return rb
}

func (rb *RequestBuilder) SetMethod(method string) *RequestBuilder {
	rb.method = method
	return rb
}

func (rb *RequestBuilder) Post() *RequestBuilder {
	rb.method = "POST"
	return rb
}

func (rb *RequestBuilder) Get() *RequestBuilder {
	rb.method = "GET"
	return rb
}

func (rb *RequestBuilder) Do(ctx context.Context) (*http.Response, error) {
	// checking required params
	if rb.url == "" {
		return nil, errors.New("url is required")
	}
	if rb.method == "" {
		return nil, errors.New("method is required")
	}

	// setting defaults
	if rb.client == nil {
		rb.client = http.DefaultClient
	}

	return rb.makeRequest(ctx)
}

func (rb *RequestBuilder) makeRequest(ctx context.Context) (*http.Response, error) {
	var (
		body io.Reader
		err  error
	)
	if len(rb.payload) > 0 {
		body, err = Marshal(rb.payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, rb.method, rb.url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range rb.params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	return rb.client.Do(req)
}
