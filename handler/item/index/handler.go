package index

import (
	"context"

	"github.com/jlevesy/readstack/repository"
)

type Handler interface {
	Handle(context.Context) (*Response, error)
}

type handler struct {
	repo repository.ItemRepository
}

func NewHandler(repo repository.ItemRepository) Handler {
	return &handler{repo}
}

func (h *handler) Handle(ctx context.Context) (*Response, error) {
	// TODO
	return &Response{}, nil
}
