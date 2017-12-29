package postgres

import (
	"context"
	"database/sql"

	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository"

	_ "github.com/lib/pq"
)

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(dbURL string) (repository.ItemRepository, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	return &itemRepository{db}, nil
}

func (i *itemRepository) Save(ctx context.Context, item *model.Item) error {
	return i.db.QueryRowContext(
		ctx,
		"INSERT INTO items(name, url) VALUES($1, $2) RETURNING id",
		item.Name,
		item.URL,
	).Scan(&item.ID)
}

func (i *itemRepository) Close() error {
	return i.db.Close()
}
