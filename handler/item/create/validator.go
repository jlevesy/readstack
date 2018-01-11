package create

import (
	"github.com/jlevesy/readstack/error"
	"github.com/jlevesy/readstack/validation"
)

type ValidatorFunc func(*Request) []*error.Violation

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
