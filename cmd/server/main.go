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
}

const (
	defaultPostgresURL     = "postgres://root:root@localhost:5672/readstack"
	defaultListenHost      = ""
	defaultListenPort      = 8080
	defaultHandlerTimemout = 200 * time.Millisecond
)

func main() {
	log.Println("Starting the server...")

	config := config{defaultPostgresURL, defaultListenHost, defaultListenPort, defaultHandlerTimemout}

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

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	item.MountRoutes(apiV1, itemRepository)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort),
			middleware.WithInMemoryTimingRecorder(
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
		),
	)
}
