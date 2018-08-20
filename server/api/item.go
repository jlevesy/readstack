package api

import (
	"context"
	"errors"

	"github.com/jlevesy/readstack/server/item"
)

type itemServer struct {
	index item.IndexHandler
}

// NewItemServer returns an instance of an ItemServer
func NewItemServer(indexHandler item.IndexHandler) ItemServer {
	return &itemServer{
		index: indexHandler,
	}
}

func (i *itemServer) Index(ctx context.Context, req *IndexRequest) (*IndexResult, error) {
	return nil, errors.New("Not implemented")
}
