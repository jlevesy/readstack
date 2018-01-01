package create

import (
	"context"

	"github.com/jlevesy/readstack/err"
	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository"
)

type Handler interface {
	Handle(ctx context.Context, req *Request) error
}

type handler struct {
	validator  ValidatorFunc
	repository repository.ItemRepository
}

func NewHandler(validator ValidatorFunc, repository repository.ItemRepository) Handler {
	return &handler{validator, repository}
}

func (h *handler) Handle(ctx context.Context, req *Request) error {
	if errs := h.validator(req); len(errs) > 0 {
		return err.NewValidationEerror(errs)
	}

	item := model.NewItem(req.Name, req.URL)

	return h.repository.Save(ctx, item)
}
