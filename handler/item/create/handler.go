package create

import (
	"context"

	readstackError "github.com/jlevesy/readstack/error"
	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository"
)

// Handler processes incoming requests and returns a result.
// It is dedicated to perform a new model.Item creation in the
// given datastore
type Handler interface {
	Handle(ctx context.Context, req *Request) error
}

type handler struct {
	validator  ValidatorFunc
	repository repository.ItemRepository
}

// NewHandler returns a Handler
func NewHandler(validator ValidatorFunc, repository repository.ItemRepository) Handler {
	return &handler{validator, repository}
}

func (h *handler) Handle(ctx context.Context, req *Request) error {
	if errs := h.validator(req); len(errs) > 0 {
		return readstackError.NewValidationError(errs)
	}

	item := model.NewItem(req.Name, req.URL)

	return h.repository.Save(ctx, item)
}
