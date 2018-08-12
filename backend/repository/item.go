package repository

import (
	"context"
	"io"

	"github.com/jlevesy/readstack/model"
)

type ItemRepository interface {
	io.Closer

	FindAll(context.Context) ([]*model.Item, error)
	Save(context.Context, *model.Item) error
}
