package mango

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// WritePlainResponse writes a plain text response with the given status code and message.
func WritePlainResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	if msg == "" {
		return
	}
	_, err := w.Write([]byte(msg))
	if err != nil {
		slog.Error("failed to write response", "err", err)
	}
}

// WriteJSONResponse writes a JSON response with the given status code and data.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.Error("failed to write JSON response", "err", err)
	}
}

/******************************************************************************/

func GetQueryParam[T any](
	r *http.Request,
	key string,
	defaultValue T,
	converter func(string) (T, error),
) (T, error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return defaultValue, nil
	}
	val, err := converter(raw)
	if err != nil {
		return defaultValue, ResponseError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("invalid query parameter: %s", key),
			Err:        err,
		}
	}
	return val, nil
}
