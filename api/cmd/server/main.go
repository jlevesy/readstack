package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jlevesy/envconfig"
	_ "github.com/lib/pq"

	"github.com/jlevesy/readstack/backend/item"
)

const (
	readstackAppName = "Readstack"

	envSeparator = "_"
)

type config struct {
	PostgresURL    string
	ListenHost     string
	ListenPort     int
	HandlerTimeout time.Duration
	WebAssetsPath  string
}

const (
	defaultPostgresURL     = "postgres://root:root@localhost:5672/readstack"
	defaultListenHost      = ""
	defaultListenPort      = 8080
	defaultHandlerTimemout = 200 * time.Millisecond
	defaultWebAssetsPath   = "dist/web"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Info("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout, defaultWebAssetsPath}

	if err := envconfig.New(readstackAppName, envSeparator).Load(&config); err != nil {
		log.Fatal(err)
	}

	logger.Info("Loaded config %v", config)

	db, err := sql.Open("postgres", config.PostgresURL)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	itemRepository, err := item.NewSQLRepository(db)

	if err != nil {
		log.Fatal(err)
	}
}
