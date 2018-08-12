package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlevesy/envconfig"

	"github.com/jlevesy/readstack/controller/errors"
	rsLogger "github.com/jlevesy/readstack/logger"
	"github.com/jlevesy/readstack/middleware"
	"github.com/jlevesy/readstack/repository/postgres"

	"github.com/jlevesy/readstack/controller/item"
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
	logger := rsLogger.NewStdLogger(
		log.New(
			os.Stdout,
			"",
			log.Ldate|log.Ltime,
		),
	)

	logger.Info("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout, defaultWebAssetsPath}

	if err := envconfig.New(readstackAppName, envSeparator).Load(&config); err != nil {
		log.Fatal(err)
	}

	logger.Info("Loaded config %v", config)

	itemRepository, err := postgres.NewItemRepository(config.PostgresURL)

	if err != nil {
		log.Fatal(err)
	}

	defer itemRepository.Close()

	errorHandler := errors.NewHandler(logger)

	r := mux.NewRouter()

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	item.MountRoutes(apiV1, itemRepository, errorHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(config.WebAssetsPath)))

	server := http.Server{
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
		IdleTimeout:       time.Second,
		Addr:              fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort),
		Handler: middleware.WithInMemoryTimingRecorder(
			middleware.Timeout(
				config.HandlerTimeout,
				middleware.RequestLogger(
					logger,
					middleware.RecordDuration(
						middleware.HandlerDuration,
						r,
					),
				),
			),
		),
	}

	logger.Fatal("%s", server.ListenAndServe())
}
