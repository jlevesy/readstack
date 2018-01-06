package err

import (
	"strings"
)

type ValidationError struct {
	errors []error
}

func NewValidationError(errors []error) *ValidationError {
	return &ValidationError{errors}
}

func (v *ValidationError) Error() string {
	msgs := make([]string, len(v.errors))

	for i, err := range v.errors {
		msgs[i] = err.Error()
	}

	return strings.Join(msgs, ",")
}
