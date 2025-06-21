package mango

import (
	"errors"
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
			w.WriteHeader(re.StatusCode)
			w.Write([]byte(re.Message))
			return
		}

		slog.Error("unexpected error in handler", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unexpected error occurred"))
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
	if e.Message == "" {
		return e.Err.Error()
	}
	return e.Message
}

func (e ResponseError) Unwrap() error {
	return e.Err
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
