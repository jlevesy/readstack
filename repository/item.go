package repository

import (
	"io"

	"github.com/jlevesy/readstack/model"
)

type ItemRepository interface {
	io.Closer

	Save(*model.Item) error
}
