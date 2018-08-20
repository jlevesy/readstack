package validation

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

// Error is an error raised when a query fails to Validate.
// It is composed by violations
type Error struct {
	Violations []*Violation
}

// NewError returns a ValidationError
func NewError(violations []*Violation) *Error {
	return &Error{violations}
}

func (v *Error) Error() string {
	msgs := make([]string, len(v.Violations))

	for i, err := range v.Violations {
		msgs[i] = err.Error()
	}

	return strings.Join(msgs, ", ")
}
