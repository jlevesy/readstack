package validation

import (
	"fmt"
	"net/url"
)

func RequireHTTPURL(attributeName, value string) error {
	req, err := url.Parse(value)

	if err != nil {
		return fmt.Errorf("Failed to parse URL for attribute %s, error is %v", attributeName, err)
	}

	if req.Scheme != "https" && req.Scheme != "http" {
		return fmt.Errorf("Unsuported URL scheme, only http and https are allowed")
	}

	return nil
}
