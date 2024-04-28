package mango

import "io"

func ReadAllString(rdr io.Reader) (string, error) {
	buf, err := io.ReadAll(rdr)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
