package item

import (
	"context"

	"github.com/jlevesy/readstack/api/validation"
)

// CreateValidatorFunc is the validation type used by create.Handler to
// validate an incoming request
type CreateValidatorFunc func(*CreateRequest) []*validation.Violation

// CreateValidator is the implementaton of the validation
func CreateValidator(req *CreateRequest) []*validation.Violation {
	res := []*validation.Violation{}

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

// CreateRequest is the create.Handler request type
// It carries all information needed in order to perform
// the create item action.
type CreateRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// CreateHandler processes incoming requests and returns a result.
// It is dedicated to perform a new model.Item creation in the
// given datastore
type CreateHandler interface {
	Handle(ctx context.Context, req *CreateRequest) error
}

type createHandler struct {
	validator  CreateValidatorFunc
	repository Repository
}

// NewCreateHandler returns a Create
func NewCreateHandler(validator CreateValidatorFunc, repository Repository) CreateHandler {
	return &createHandler{validator, repository}
}

func (h *createHandler) Handle(ctx context.Context, req *CreateRequest) error {
	if errs := h.validator(req); len(errs) > 0 {
		return validation.NewError(errs)
	}

	item := New(req.Name, req.URL)

	return h.repository.Create(ctx, item)
}
