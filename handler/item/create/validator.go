package create

import (
	"github.com/jlevesy/readstack/error"
	"github.com/jlevesy/readstack/validation"
)

// ValidatorFunc is the validation type used by create.Handler to
// validate an incoming request
type ValidatorFunc func(*Request) []*error.Violation

// Validator is the implementaton of the validation
func Validator(req *Request) []*error.Violation {
	res := []*error.Violation{}

	if v := validation.RequireNotBlank("Name", req.Name); v != nil {
		res = append(res, v)
	}

	if v := validation.RequireNotBlank("URL", req.URL); v != nil {
		res = append(res, v)
	}

	if v := validation.RequireHTTPURL("URL", req.URL); v != nil {
		res = append(res, v)
	}

	return res
}
