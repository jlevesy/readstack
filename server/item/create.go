package item

import (
	"context"

	"github.com/jlevesy/readstack/server/validation"
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
	Name string
	URL  string
}

// CreateResponse carries the created item
type CreateResponse struct {
	CreatedItem *Model
}

// CreateHandler processes incoming requests and returns a result.
// It is dedicated to perform a new model.Item creation in the
// given datastore
type CreateHandler interface {
	Handle(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
}

type createHandler struct {
	validator  CreateValidatorFunc
	repository Repository
}

// NewCreateHandler returns a Create
func NewCreateHandler(validator CreateValidatorFunc, repository Repository) CreateHandler {
	return &createHandler{validator, repository}
}

func (h *createHandler) Handle(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if errs := h.validator(req); len(errs) > 0 {
		return nil, validation.NewError(errs)
	}

	item := New(req.Name, req.URL)

	err := h.repository.Create(ctx, item)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{item}, nil
}
