package game

import "errors"

var (
	// ErrValidation thrown when a validation error occurs
	ErrValidation = errors.New("validation")

	// ErrRegex thrown when a regex error occurs
	ErrRegex = errors.New("regex")
)

