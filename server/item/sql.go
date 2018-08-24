package item

import (
	"context"
	"database/sql"
)

type sqlRepository struct {
	create  *sql.Stmt
	findAll *sql.Stmt
	delete  *sql.Stmt
}

// NewSQLRepository returns an SQL repository
func NewSQLRepository(db *sql.DB) (Repository, error) {
	createStmt, err := db.Prepare("INSERT INTO items(name, url) VALUES($1, $2) RETURNING id")

	if err != nil {
		return nil, err
	}

	findAllStmt, err := db.Prepare("SELECT * FROM items")

	if err != nil {
		return nil, err
	}

	deleteStmt, err := db.Prepare("DELETE FROM items where id = $1")

	if err != nil {
		return nil, err
	}

	return &sqlRepository{createStmt, findAllStmt, deleteStmt}, nil
}

func (i *sqlRepository) Create(ctx context.Context, item *Model) error {
	return i.create.QueryRowContext(ctx, item.Name, item.URL).Scan(&item.id)
}

func (i *sqlRepository) FindAll(ctx context.Context) ([]*Model, error) {
	result := []*Model{}
	rows, err := i.findAll.QueryContext(ctx)

	if err != nil {
		return result, err
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return result, err
	}

	for rows.Next() {
		item := Model{}

		if err := rows.Scan(&item.id, &item.Name, &item.URL); err != nil {
			return result, err
		}

		result = append(result, &item)
	}

	return result, nil
}

func (i *sqlRepository) Delete(ctx context.Context, item *Model) error {
	res, err := i.delete.ExecContext(ctx, item.GetID())

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return ErrItemNotFound
	}

	return nil
}
