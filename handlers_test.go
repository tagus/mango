package mango

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapErrorHandler(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:           "no error",
			expectedStatus: http.StatusOK,
		},
		{
			name:            "bad request error",
			err:             BadRequestError("bad request"),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "bad request",
		},
		{
			name:            "internal server error",
			err:             InternalServerError("error occurred", fmt.Errorf("some internal error")),
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: "error occurred",
		},
		{
			name:            "unexpected error",
			err:             fmt.Errorf("unexpected error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: "unexpected error occurred",
		},
		{
			name:            "response error with message",
			err:             ResponseError{StatusCode: http.StatusForbidden, Message: "forbidden"},
			expectedStatus:  http.StatusForbidden,
			expectedMessage: "forbidden",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := WrapErrorHandler(func(w http.ResponseWriter, r *http.Request) error {
				return test.err
			})

			rr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			handler(rr, req)

			res := rr.Result()
			assert.Equal(t, test.expectedStatus, res.StatusCode)
			assert.Equal(t, test.expectedMessage, rr.Body.String())
		})
	}
}

func TestResponseError_MarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		responseErr  ResponseError
		expectedJSON string
	}{
		{
			name: "with underlying error",
			responseErr: ResponseError{
				StatusCode: http.StatusBadRequest,
				Message:    "bad request",
				Err:        fmt.Errorf("underlying error"),
			},
			expectedJSON: `{"message":"bad request","error":"underlying error"}`,
		},
		{
			name: "without underlying error",
			responseErr: ResponseError{
				StatusCode: http.StatusNotFound,
				Message:    "not found",
			},
			expectedJSON: `{"message":"not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.responseErr.MarshalJSON()
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expectedJSON, string(data))
		})
	}
}
