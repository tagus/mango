package mango

import "fmt"

var (
	ErrValidation = fmt.Errorf("validation error")
)

type validationError struct {
	msg string
}

// ValidationError represents a bad input condition that should typically be
// surfaced to callers as a 400-style client error.
func ValidationError(msg string) error {
	return &validationError{msg: msg}
}

func (e *validationError) Error() string {
	return e.msg
}

func (e *validationError) Is(target error) bool {
	return target == ErrValidation
}
