package validation

import (
	"fmt"
	"net/url"

	"github.com/jlevesy/readstack/handler/errors"
)

const (
	reasonParseFailed      = "Failed to parse URL, reason is :%s"
	reasonUnsuportedScheme = "Unsuported URL scheme, only http and https are allowed"
)

func RequireHTTPURL(attributeName, value string) *errors.Violation {
	req, err := url.Parse(value)

	if err != nil {
		return &errors.Violation{
			Name:   attributeName,
			Reason: fmt.Sprintf(reasonParseFailed, err),
		}
	}

	if req.Scheme != "https" && req.Scheme != "http" {
		return &errors.Violation{
			Name:   attributeName,
			Reason: reasonUnsuportedScheme,
		}
	}

	return nil
}
