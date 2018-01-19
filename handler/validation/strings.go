package validation

import (
	"strings"

	"github.com/jlevesy/readstack/handler/errors"
)

const (
	reasonNotBlank = "Should not be blank"
)

func RequireNotBlank(attributeName, value string) *errors.Violation {
	if strings.TrimSpace(value) == "" {
		return &errors.Violation{
			Name:   attributeName,
			Reason: reasonNotBlank,
		}
	}

	return nil
}
