package mango

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

func ShortID() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
