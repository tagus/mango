package mango

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func WrapErrorHandler(h ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		var re ResponseError
		if errors.As(err, &re) {
			if re.Err != nil {
				slog.Error("unexpect error in handler", "err", re.Err)
			}
			accept := r.Header.Get("Accept")
			if accept == "application/json" {
				WriteJSONResponse(w, re.StatusCode, re)
			} else {
				WritePlainResponse(w, re.StatusCode, re.Message)
			}
			return
		}

		slog.Error("unexpected error in handler", "err", err)
		WritePlainResponse(w, http.StatusInternalServerError, "unexpected error occurred")
	}
}

/******************************************************************************/

type ResponseError struct {
	// the standard HTTP response status code
	StatusCode int
	// the response message sent to send back to the client
	Message string
	// the underlying error
	Err error
}

func (e ResponseError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	if e.Message == "" {
		return e.Err.Error()
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e ResponseError) Unwrap() error {
	return e.Err
}

func (e ResponseError) MarshalJSON() ([]byte, error) {
	type resp struct {
		Message string `json:"message"`
		Error   string `json:"error,omitempty"`
	}
	res := resp{
		Message: e.Message,
	}
	if e.Err != nil {
		res.Error = e.Err.Error()
	}
	return json.Marshal(res)
}

func BadRequestError(msg string) ResponseError {
	return ResponseError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NotFoundError(msg string) ResponseError {
	return ResponseError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func UnauthorizedError(msg string) ResponseError {
	return ResponseError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
	}
}

func InternalServerError(msg string, err error) ResponseError {
	return ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
		Err:        err,
	}
}
