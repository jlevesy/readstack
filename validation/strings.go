package validation

import (
	"strings"
)

const (
	reasonNotBlank = "Should not be blank"
)

func RequireNotBlank(attributeName, value string) *Violation {
	if strings.TrimSpace(value) == "" {
		return &Violation{
			Name:   attributeName,
			Reason: reasonNotBlank,
		}
	}

	return nil
}
