package repository

import (
	"github.com/jlevesy/readstack/model"
)

type Article interface {
	Save(*model.Article) (*model.Article, error)
}
