package create

import (
	"github.com/jlevesy/readstack/validation"
)

type ValidatorFunc func(*Request) []error

func Validator(req *Request) []error {
	res := []error{}

	if err := validation.RequireNotBlank("Name", req.Name); err != nil {
		res = append(res, err)
	}

	if err := validation.RequireNotBlank("URL", req.URL); err != nil {
		res = append(res, err)
	}

	if err := validation.RequireHTTPURL("URL", req.URL); err != nil {
		res = append(res, err)
	}

	return res
}
