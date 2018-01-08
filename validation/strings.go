package validation

import (
	"fmt"
	"strings"
)

func RequireNotBlank(attributeName, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("field %s not supposed to be blank", attributeName)
	}

	return nil
}
