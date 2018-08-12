package index

import (
	"context"

	"github.com/jlevesy/readstack/handler/item/index"
)

type HandlerStub struct {
	OnHandle func(ctx context.Context) (*index.Response, error)
}

func (h *HandlerStub) Handle(ctx context.Context) (*index.Response, error) {
	return h.OnHandle(ctx)
}
