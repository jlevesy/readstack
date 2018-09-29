package api

import (
	"context"

	"github.com/jlevesy/readstack/item"
)

type protobufServer struct {
	index  item.IndexHandler
	create item.CreateHandler
	delete item.DeleteHandler
}

// NewItemServer returns an instance of an ItemServer
func NewItemServer(indexHandler item.IndexHandler, create item.CreateHandler, delete item.DeleteHandler) ItemServer {
	return &protobufServer{
		index:  indexHandler,
		create: create,
		delete: delete,
	}
}

func (i *protobufServer) Index(ctx context.Context, _ *IndexRequest) (*IndexResult, error) {
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

func (i *protobufServer) Create(ctx context.Context, req *CreateRequest) (*CreateResult, error) {
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

func (i *protobufServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResult, error) {
	err := i.delete.Handle(ctx, &item.DeleteRequest{ID: req.GetId()})

	if err != nil {
		return nil, err
	}

	return &DeleteResult{}, nil
}

func toIndexItem(item *item.Model) *IndexItem {
	return &IndexItem{
		Id:   item.GetID(),
		Name: item.Name,
		Url:  item.URL,
	}
}
