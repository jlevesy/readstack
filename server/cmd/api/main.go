package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jlevesy/envconfig"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/jlevesy/readstack/server/api"
	"github.com/jlevesy/readstack/server/item"
)

const (
	readstackAppName = "Readstack"

	envSeparator = "_"
)

type config struct {
	PostgresURL string
	ListenHost  string
	ListenPort  int
}

func (c *config) listenURL() string {
	return fmt.Sprintf("%s:%d", c.ListenHost, c.ListenPort)
}

const (
	defaultPostgresURL = "postgres://root:root@localhost:5672/readstack"
	defaultListenHost  = ""
	defaultListenPort  = 8080
)

func main() {
	logger := logrus.New()
	entry := logrus.NewEntry(logger)

	logger.Info("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort}

	if err := envconfig.New(readstackAppName, envSeparator).Load(&config); err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	logger.Infof("Loaded config %v", config)

	lis, err := net.Listen("tcp", config.listenURL())

	if err != nil {
		logger.WithError(err).Fatal("Failed to open listening socket")
	}

	db, err := sql.Open("postgres", config.PostgresURL)

	if err != nil {
		logger.WithError(err).Fatal("Failed to open database connection")
	}

	itemRepository, err := item.NewSQLRepository(db)

	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize the item repository")
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(entry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(entry),
		),
	)

	api.RegisterItemServer(
		s,
		api.NewItemServer(
			item.NewIndexHandler(itemRepository),
			item.NewCreateHandler(item.CreateValidator, itemRepository),
			item.NewDeleteHandler(itemRepository),
		),
	)

	if err := s.Serve(lis); err != nil {
		logger.WithError(err).Error("Server exited with an error")
	}
}
