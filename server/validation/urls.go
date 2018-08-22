package validation

import (
	"fmt"
	"net/url"
)

const (
	reasonParseFailed      = "Failed to parse URL, reason is :%s"
	reasonUnsuportedScheme = "Unsuported URL scheme, only http and https are allowed"
)

func RequireHTTPURL(attributeName, value string) *Violation {
	req, err := url.Parse(value)

	if err != nil {
		return &Violation{
			Name:   attributeName,
			Reason: fmt.Sprintf(reasonParseFailed, err),
		}
	}

	if req.Scheme != "https" && req.Scheme != "http" {
		return &Violation{
			Name:   attributeName,
			Reason: reasonUnsuportedScheme,
		}
	}

	return nil
}
