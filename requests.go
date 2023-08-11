package mango

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type RequestBuilder struct {
	url     string
	method  string
	payload any
	params  map[string]string
	client  *http.Client
}

func NewRequest() *RequestBuilder {
	return &RequestBuilder{}
}

func GetRequest() *RequestBuilder {
	return &RequestBuilder{method: "GET"}
}

func PostRequest() *RequestBuilder {
	return &RequestBuilder{method: "POST"}
}

func (rb *RequestBuilder) SetUrl(url string) *RequestBuilder {
	rb.url = url
	return rb
}

func (rb *RequestBuilder) SetPayload(payload any) *RequestBuilder {
	rb.payload = payload
	return rb
}

func (rb *RequestBuilder) SetParams(params map[string]string) *RequestBuilder {
	rb.params = params
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
	var body io.Reader
	if rb.payload != nil {
		_body, err := Marshal(rb.payload)
		if err != nil {
			return nil, err
		}
		body = _body
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
