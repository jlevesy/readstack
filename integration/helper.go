package integration

import (
	"database/sql"
	"net"
	"testing"

	// import the postgres driver
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"google.golang.org/grpc"

	"github.com/jlevesy/readstack/api"
	"github.com/jlevesy/readstack/item"
)

const (
	dbStr         = "host=db user=readstack password=notsecret dbname=readstack_test sslmode=disable"
	migrationPath = "../db/migrations"
)

type testContext struct {
	db         *sql.DB
	migrations *migrate.FileMigrationSource
	server     *grpc.Server

	clientConn *grpc.ClientConn

	ItemClient api.ItemClient
}

func setup(t *testing.T) *testContext {
	var (
		tc  testContext
		err error
	)

	tc.db, err = sql.Open("postgres", dbStr)

	if err != nil {
		t.Fatal("Failed to connect to the SQL server", err)
	}

	tc.migrations = &migrate.FileMigrationSource{
		Dir: "../db/migrations",
	}

	if _, err := migrate.Exec(tc.db, "postgres", tc.migrations, migrate.Up); err != nil {
		t.Fatal("Failed to execute migrations", err)
	}

	itemRepository, err := item.NewSQLRepository(tc.db)

	if err != nil {
		t.Fatal("Failed to setup the SQL repository", err)
	}

	lis, err := net.Listen("tcp", ":60811")

	if err != nil {
		t.Fatal("Failed to bind to the test port", err)
	}

	tc.server = grpc.NewServer()

	api.RegisterItemServer(
		tc.server,
		api.NewItemServer(
			item.NewIndexHandler(itemRepository),
			item.NewCreateHandler(item.CreateValidator, itemRepository),
			item.NewDeleteHandler(itemRepository),
		),
	)

	go tc.server.Serve(lis)

	tc.clientConn, err = grpc.Dial("localhost:60811", grpc.WithInsecure())

	if err != nil {
		t.Fatal("Failed to connect to the testServer", err)
	}

	tc.ItemClient = api.NewItemClient(tc.clientConn)

	return &tc
}

func (t *testContext) TearDown() {
	defer t.db.Close()
	defer migrate.Exec(t.db, "postgres", t.migrations, migrate.Down)
	defer t.server.GracefulStop()
	defer t.clientConn.Close()
}
