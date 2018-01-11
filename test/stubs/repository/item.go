package item

import (
	"context"

	"github.com/jlevesy/readstack/model"
)

type ItemRepositoryStub struct {
	CloseStub func()

	FindAllStub func(context.Context) ([]*model.Item, error)
	SaveStub    func(context.Context, *model.Item) error
}
