package index

import (
	"context"

	"github.com/jlevesy/readstack/repository"
)

// Handler processes incoming requests and returns a result.
// It is dedicated to find all items currently stored in given
// data store.
type Handler interface {
	Handle(context.Context) (*Response, error)
}

type handler struct {
	repo repository.ItemRepository
}

// NewHandler returns a Handler
func NewHandler(repo repository.ItemRepository) Handler {
	return &handler{repo}
}

func (h *handler) Handle(ctx context.Context) (*Response, error) {
	res, err := h.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return &Response{res}, nil
}
