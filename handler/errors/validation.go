package errors

import (
	"fmt"

	"strings"
)

// Violation represents a rule violation by a query
type Violation struct {
	Name   string
	Reason string
}

func (v *Violation) Error() string {
	return fmt.Sprintf("%s : %s", v.Name, v.Reason)
}

// ValidationError is an error raised when a query fails to Validate.
// It is composed by violations
type ValidationError struct {
	Violations []*Violation
}

// NewValidationError returns a ValidationError
func NewValidationError(violations []*Violation) *ValidationError {
	return &ValidationError{violations}
}

func (v *ValidationError) Error() string {
	msgs := make([]string, len(v.Violations))

	for i, err := range v.Violations {
		msgs[i] = err.Error()
	}

	return strings.Join(msgs, ", ")
}
