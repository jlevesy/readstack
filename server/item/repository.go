package item

import (
	"context"
)

type Repository interface {
	FindAll(context.Context) ([]*Model, error)
	Create(context.Context, *Model) error
	Delete(context.Context, *Model) error
}
