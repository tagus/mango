package mango

import (
	"strconv"
	"strings"
)

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

// ParseList converts the given comma separated string into a typed slice.
func ParseList[T any](s string, converter func(string) (T, error)) ([]T, error) {
	if s == "" {
		return []T{}, nil
	}
	parts := strings.Split(s, ",")
	results := make([]T, 0, len(parts))
	for _, part := range parts {
		cleaned := strings.TrimSpace(part)
		if cleaned == "" {
			continue
		}
		value, err := converter(cleaned)
		if err != nil {
			return nil, err
		}
		results = append(results, value)
	}
	return results, nil
}

// ParseStringList is a shortcut for parsing a string slice
func ParseStringList(s string) ([]string, error) {
	return ParseList(s, ParseString)
}
