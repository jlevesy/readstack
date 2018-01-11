package validation

import (
	"fmt"
	"net/url"

	"github.com/jlevesy/readstack/error"
)

const (
	reasonParseFailed      = "Failed to parse URL, reason is :%s"
	reasonUnsuportedScheme = "Unsuported URL scheme, only http and https are allowed"
)

func RequireHTTPURL(attributeName, value string) *error.Violation {
	req, err := url.Parse(value)

	if err != nil {
		return &error.Violation{
			Name:   attributeName,
			Reason: fmt.Sprintf(reasonParseFailed, err),
		}
	}

	if req.Scheme != "https" && req.Scheme != "http" {
		return &error.Violation{
			Name:   attributeName,
			Reason: reasonUnsuportedScheme,
		}
	}

	return nil
}
