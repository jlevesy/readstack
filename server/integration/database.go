package integration

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
)

const (
	dbStr = "host=db user=readstack password=notsecret dbname=readstack_test sslmode=disable"
)

func SetupDB() (*sql.DB, func(), error) {
	db, err := sql.Open("postgres", dbStr)

	if err != nil {
		return nil, nil, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "../db/migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return nil, nil, err
	}

	done := func() {
		defer db.Close()
		migrate.Exec(db, "postgres", migrations, migrate.Down)
	}

	return db, done, nil
}
