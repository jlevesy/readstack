package create

import (
	"context"

	"github.com/jlevesy/readstack/handler/item/create"
)

// HandlerStub is a stub to implement create.Handler interface
// in tests
type HandlerStub struct {
	OnHandle func(ctx context.Context, req *create.Request) error
}

func (h *HandlerStub) Handle(ctx context.Context, req *create.Request) error {
	return h.OnHandle(ctx, req)
}
