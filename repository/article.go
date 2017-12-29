package repository

import (
	"github.com/jlevesy/readstack/model"
)

type ArticleRepository interface {
	Save(*model.Article) error
}
