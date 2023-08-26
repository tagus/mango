package mango

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

// TokenFromFile will attempt to parse an oauth token
// at the provided file path
func TokenFromFile(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, err
	}

	return tok, nil
}
