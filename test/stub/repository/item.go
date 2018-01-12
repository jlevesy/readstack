package repository

import (
	"context"

	"github.com/jlevesy/readstack/model"
)

type ItemRepositoryStub struct {
	OnClose func() error

	OnFindAll func(context.Context) ([]*model.Item, error)
	OnSave    func(context.Context, *model.Item) error
}

func (i *ItemRepositoryStub) Close() error {
	return i.OnClose()
}

func (i *ItemRepositoryStub) FindAll(ctx context.Context) ([]*model.Item, error) {
	return i.OnFindAll(ctx)
}

func (i *ItemRepositoryStub) Save(ctx context.Context, item *model.Item) error {
	return i.OnSave(ctx, item)
}
