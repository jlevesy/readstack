package postgres

import (
	"database/sql"

	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository"

	_ "github.com/lib/pq"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(dbURL string) (repository.ArticleRepository, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	return &articleRepository{db}, nil
}

func (a *articleRepository) Save(article *model.Article) error {
	return a.db.QueryRow(
		`INSERT INTO articles(name, url) VALUES (?, ?) RETURNING id`,
		article.Name,
		article.URL,
	).Scan(article.ID)
}

func (a *articleRepository) Close() error {
	return a.db.Close()
}
