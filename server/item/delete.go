package item

import (
	"context"
)

// DeleteRequest is the delete.Handler request type
// It carries all information needed in order to perform
// the delete item action.
type DeleteRequest struct {
	ID int64
}

// DeleteHandler processes incoming requests and returns a result.
// It is dedicated to perform a new model.Item creation in the
// given datastore
type DeleteHandler interface {
	Handle(ctx context.Context, req *DeleteRequest) error
}

type deleteHandler struct {
	repository Repository
}

// NewDeleteHandler returns a new DeleteHandler action
func NewDeleteHandler(repository Repository) DeleteHandler {
	return &deleteHandler{repository}
}

func (h *deleteHandler) Handle(ctx context.Context, req *DeleteRequest) error {
	return h.repository.Delete(ctx, &Model{id: req.ID})
}
