package api

import (
	"context"

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

func (i *itemServer) Index(ctx context.Context, _ *IndexRequest) (*IndexResult, error) {
	result, err := i.index.Handle(ctx)

	if err != nil {
		return nil, err
	}

	items := make([]*IndexItem, 0, len(result.Items))

	for _, item := range result.Items {
		items = append(items, toIndexItem(item))
	}

	return &IndexResult{Items: items}, nil
}

func toIndexItem(item *item.Model) *IndexItem {
	return &IndexItem{
		Id:   item.GetID(),
		Name: item.Name,
		Url:  item.URL,
	}
}
