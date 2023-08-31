package mango

import (
	"hash/fnv"
	"strconv"
)

// HashStrings computes the hash of the given strings by hashing the strings
// together into an int representation and then converting that to a string
// to keep the resulting string small
func HashStrings(vals ...string) (string, error) {
	hash := fnv.New32a()
	for _, val := range vals {
		_, err := hash.Write([]byte(val))
		if err != nil {
			return "", err
		}
	}
	res := hash.Sum32()
	return strconv.Itoa(int(res)), nil
}
