package mango

import (
	"bytes"
	"encoding/json"
	"io"
)

func Unmarshal[T any](rdr io.Reader) (*T, error) {
	decoded := new(T)
	if err := json.NewDecoder(rdr).Decode(&decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}

func UnmarshalFromString[T any](str string) (*T, error) {
	buf := bytes.NewBufferString(str)
	return Unmarshal[T](buf)
}

func Marshal(payload any) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(payload); err != nil {
		return nil, err
	}
	return buf, nil
}

func MarshalToString(payload any) (string, error) {
	rdr, err := Marshal(payload)
	if err != nil {
		return "", nil
	}
	return ReadAllString(rdr)
}
