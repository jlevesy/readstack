package error

import (
	"fmt"

	"strings"
)

type Violation struct {
	Name   string
	Reason string
}

func (v *Violation) Error() string {
	return fmt.Sprintf("%s : %s", v.Name, v.Reason)
}

type ValidationError struct {
	Violations []*Violation
}

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
