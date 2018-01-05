package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jlevesy/envconfig"

	"github.com/jlevesy/readstack/middleware"
	"github.com/jlevesy/readstack/repository/postgres"

	"github.com/jlevesy/readstack/controller/item"
	createItem "github.com/jlevesy/readstack/handler/item/create"
)

const (
	readstackAppName = "READSTACK"
	defaultSeparator = "_"
)

type config struct {
	PostgresURL string
}

func main() {
	log.Println("Starting the server...")

	config := config{}

	if err := envconfig.New(readstackAppName, defaultSeparator).Load(&config); err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded config %v", config)

	itemRepository, err := postgres.NewItemRepository(config.PostgresURL)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle(
		"/item",
		middleware.WithInMemoryTimingRecorder(
			middleware.Timeout(
				200*time.Millisecond,
				middleware.RequestLogger(
					middleware.RecordDuration(
						middleware.HandlerDuration,
						middleware.Post(
							item.NewCreateController(
								createItem.NewHandler(
									createItem.Validator,
									itemRepository,
								),
							),
						),
					),
				),
			),
		),
	)

	http.ListenAndServe(":8080", nil)
}
