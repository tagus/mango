package mango

import (
	"os"
	"strings"
)

// GetEnvString retrieves the env variable for the given key and parses it as a string slice.
func GetEnvStringSlice(key string, defaultValue []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return defaultValue
	}
	args := strings.Split(val, ",")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	return args
}
