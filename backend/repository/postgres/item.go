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
	stmt, err := i.db.Prepare("INSERT INTO items(name, url) VALUES($1, $2) RETURNING id")

	if err != nil {
		return err
	}

	defer stmt.Close()

	return stmt.QueryRowContext(ctx, item.Name, item.URL).Scan(&item.ID)
}

func (i *itemRepository) FindAll(ctx context.Context) ([]*model.Item, error) {
	stmt, err := i.db.Prepare("SELECT * FROM items")

	if err != nil {
		return []*model.Item{}, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)

	if err != nil {
		return []*model.Item{}, err
	}

	defer rows.Close()

	result := []*model.Item{}

	for rows.Next() {
		var (
			id   int64
			name string
			URL  string
		)

		if err := rows.Scan(&id, &name, &URL); err != nil {
			return []*model.Item{}, err
		}

		result = append(result, &model.Item{id, name, URL})
	}

	if err := rows.Err(); err != nil {
		return []*model.Item{}, err
	}

	return result, nil
}

func (i *itemRepository) Close() error {
	return i.db.Close()
}
