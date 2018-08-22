package api

import (
	"context"

	"github.com/jlevesy/readstack/server/item"
)

type itemServer struct {
	index  item.IndexHandler
	create item.CreateHandler
}

// NewItemServer returns an instance of an ItemServer
func NewItemServer(indexHandler item.IndexHandler, create item.CreateHandler) ItemServer {
	return &itemServer{
		index:  indexHandler,
		create: create,
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

func (i *itemServer) Create(ctx context.Context, req *CreateRequest) (*CreateResult, error) {
	result, err := i.create.Handle(ctx, &item.CreateRequest{Name: req.GetName(), URL: req.GetUrl()})

	if err != nil {
		return nil, err
	}

	return &CreateResult{
		Id:   result.CreatedItem.GetID(),
		Name: result.CreatedItem.Name,
		Url:  result.CreatedItem.URL,
	}, nil

}

func toIndexItem(item *item.Model) *IndexItem {
	return &IndexItem{
		Id:   item.GetID(),
		Name: item.Name,
		Url:  item.URL,
	}
}
