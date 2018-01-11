package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlevesy/envconfig"

	"github.com/jlevesy/readstack/middleware"
	"github.com/jlevesy/readstack/repository/postgres"

	"github.com/jlevesy/readstack/controller/item"
)

const (
	readstackAppName = "READSTACK"
	defaultSeparator = "_"
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
	log.Println("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout, defaultWebAssetsPath}

	if err := envconfig.New(readstackAppName, defaultSeparator).Load(&config); err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded config %v", config)

	itemRepository, err := postgres.NewItemRepository(config.PostgresURL)

	if err != nil {
		log.Fatal(err)
	}

	defer itemRepository.Close()

	r := mux.NewRouter()

	r.Path("/").Handler(http.FileServer(http.Dir(config.WebAssetsPath)))

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	item.MountRoutes(apiV1, itemRepository)

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
					middleware.RecordDuration(
						middleware.HandlerDuration,
						r,
					),
				),
			),
		),
	}

	log.Fatal(server.ListenAndServe())
}
