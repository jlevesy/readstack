package validation

import (
	"strings"

	"github.com/jlevesy/readstack/error"
)

const (
	reasonNotBlank = "Should not be blank"
)

func RequireNotBlank(attributeName, value string) *error.Violation {
	if strings.TrimSpace(value) == "" {
		return &error.Violation{
			Name:   attributeName,
			Reason: reasonNotBlank,
		}
	}

	return nil
}
