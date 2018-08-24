package item

import (
	"context"
	"errors"
)

var (
	//ErrItemNotFound means required item is not found
	ErrItemNotFound = errors.New("Item not found")
)

type Repository interface {
	FindAll(context.Context) ([]*Model, error)
	Create(context.Context, *Model) error
	Delete(context.Context, *Model) error
}
