package item

import (
	"context"
)

type RepositoryStub struct {
	OnFindAll func(context.Context) ([]*Model, error)
	OnCreate  func(context.Context, *Model) error
}

func (i *RepositoryStub) FindAll(ctx context.Context) ([]*Model, error) {
	return i.OnFindAll(ctx)
}

func (i *RepositoryStub) Create(ctx context.Context, item *Model) error {
	return i.OnCreate(ctx, item)
}
