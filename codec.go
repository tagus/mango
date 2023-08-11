package mango

import (
	"bytes"
	"encoding/json"
	"io"
)

func Unmarshal[T any](rdr io.ReadCloser) (*T, error) {
	decoded := new(T)
	if err := json.NewDecoder(rdr).Decode(&decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}

func Marshal(payload any) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(payload); err != nil {
		return nil, err
	}
	return buf, nil
}
