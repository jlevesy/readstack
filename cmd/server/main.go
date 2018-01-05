package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jlevesy/readstack/middleware"
	"github.com/jlevesy/readstack/repository/postgres"

	"github.com/jlevesy/readstack/controller/item"
	createItem "github.com/jlevesy/readstack/handler/item/create"
)

const (
	postgresURL = "postgres://readstack:notsecret@db:5432/readstack?sslmode=disable"
)

func main() {
	log.Println("Starting the server...")

	itemRepository, err := postgres.NewItemRepository(postgresURL)

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
