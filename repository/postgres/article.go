package postgres

import (
	"database/sql"

	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository"

	_ "github.com/lib/pq"
)

type articleRepository struct {
	db sql.DB
}

func NewArticleRepository(dbURL, string) (repository.ArticleRepository, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	return &articleRepository{db}, nil
}

func (a *articleRepository) Save(a *model.Article) error {
	return r.db.QueryRow(
		`INSERT INTO articles(name, url) VALUES (?, ?) RETURNING id`, a.Name, a.URL,
	).Scan(a.ID)
}
