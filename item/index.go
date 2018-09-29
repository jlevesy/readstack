package item

import (
	"context"
)

// IndexResponse represents IndexHandler result
type IndexResponse struct {
	Items []*Model
}

// IndexHandler find all items currently stored
type IndexHandler interface {
	Handle(context.Context) (*IndexResponse, error)
}

type indexHandler struct {
	repo Repository
}

// NewIndexHandler returns a new indexHandler instance
func NewIndexHandler(repo Repository) IndexHandler {
	return &indexHandler{repo}
}
func (i *indexHandler) Handle(ctx context.Context) (*IndexResponse, error) {
	res, err := i.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return &IndexResponse{res}, nil
}
