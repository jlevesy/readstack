package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/jlevesy/envconfig"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/jlevesy/readstack/server/api"
	"github.com/jlevesy/readstack/server/item"
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

func (c *config) listenURL() string {
	return fmt.Sprintf("%s:%d", c.ListenHost, c.ListenPort)
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

	logger.Println("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout, defaultWebAssetsPath}

	if err := envconfig.New(readstackAppName, envSeparator).Load(&config); err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Loaded config %v", config)

	lis, err := net.Listen("tcp", config.listenURL())

	if err != nil {
		logger.Fatal(err)
	}

	db, err := sql.Open("postgres", config.PostgresURL)

	if err != nil {
		logger.Fatal(err)
	}

	itemRepository, err := item.NewSQLRepository(db)

	if err != nil {
		logger.Fatal(err)
	}

	s := grpc.NewServer()

	api.RegisterItemServer(
		s,
		api.NewItemServer(
			item.NewIndexHandler(itemRepository),
			item.NewCreateHandler(item.CreateValidator, itemRepository),
			item.NewDeleteHandler(itemRepository),
		),
	)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
