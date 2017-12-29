package repository

import (
	"context"
	"io"

	"github.com/jlevesy/readstack/model"
)

type ItemRepository interface {
	io.Closer

	Save(context.Context, *model.Item) error
}
