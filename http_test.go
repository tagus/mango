package mango

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePlainResponse(t *testing.T) {
	tests := []struct {
		name               string
		statusCode         int
		message            string
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name:               "simple message",
			statusCode:         http.StatusOK,
			message:            "sup",
			expectedStatusCode: http.StatusOK,
			expectedMessage:    "sup",
		},
		{
			name:               "not found",
			statusCode:         http.StatusNotFound,
			message:            "not found",
			expectedStatusCode: http.StatusNotFound,
			expectedMessage:    "not found",
		},
		{
			name:               "empty message",
			statusCode:         http.StatusNoContent,
			message:            "",
			expectedStatusCode: http.StatusNoContent,
			expectedMessage:    "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, _ *http.Request) {
				WritePlainResponse(w, test.statusCode, test.message)
			}

			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			handler(rr, req)

			res := rr.Result()
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)
			assert.Equal(t, test.expectedMessage, rr.Body.String())
			assert.Equal(t, "text/plain; charset=utf-8", res.Header.Get("Content-Type"))
		})
	}
}

func TestWriteJSONResponse(t *testing.T) {
	tests := []struct {
		name               string
		statusCode         int
		data               any
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "simple JSON",
			statusCode:         http.StatusOK,
			data:               map[string]string{"message": "hello"},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"message":"hello"}` + "\n",
		},
		{
			name:               "empty data",
			statusCode:         http.StatusNoContent,
			data:               nil,
			expectedStatusCode: http.StatusNoContent,
			expectedBody:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, _ *http.Request) {
				WriteJSONResponse(w, test.statusCode, test.data)
			}

			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			handler(rr, req)

			res := rr.Result()
			assert.Equal(t, test.expectedStatusCode, res.StatusCode)
			assert.Equal(t, test.expectedBody, rr.Body.String())
			assert.Equal(t, "application/json; charset=utf-8", res.Header.Get("Content-Type"))
		})
	}
}

func TestGetQueryParam(t *testing.T) {
	tests := []struct {
		name          string
		query         string
		key           string
		defaultValue  string
		expectedValue string
		expectErr     bool
		converter     func(string) (string, error)
	}{
		{
			name:          "param exists",
			query:         "?foo=bar",
			key:           "foo",
			defaultValue:  "default",
			expectedValue: "bar",
			converter: func(s string) (string, error) {
				return s, nil
			},
		},
		{
			name:          "param does not exist",
			query:         "?baz=qux",
			key:           "foo",
			defaultValue:  "default",
			expectedValue: "default",
			converter: func(s string) (string, error) {
				return s, nil
			},
		},
		{
			name:          "empty query",
			query:         "",
			key:           "foo",
			defaultValue:  "default",
			expectedValue: "default",
			converter: func(s string) (string, error) {
				return s, nil
			},
		},
		{
			name:          "invalid value",
			query:         "?foo=invalid",
			key:           "foo",
			defaultValue:  "default",
			expectedValue: "default",
			expectErr:     true,
			converter: func(s string) (string, error) {
				return "", fmt.Errorf("invalid value")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/"+test.query, nil)
			value, err := GetQueryParam(req, test.key, test.defaultValue, test.converter)
			if test.expectErr {
				assert.Error(t, err, "expected an error but got none")
				return
			}
			assert.NoError(t, err, "expected no error but got one")
			assert.Equal(t, test.expectedValue, value)
		})
	}
}
