package mango

import "strconv"

func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// ParseString is an identity function for strings to adapt to the expected interface.
func ParseString(s string) (string, error) {
	return s, nil
}
