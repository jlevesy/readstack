package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jlevesy/readstack/repository/postgres"

	"github.com/jlevesy/readstack/controller/item"
	createItem "github.com/jlevesy/readstack/handler/item/create"
)

const (
	postgresURL = "postgres://readstack:notsecret@db:5432/readstack?sslmode=disable"
)

func post(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			if req.Method != http.MethodPost {
				http.NotFound(w, req)
				return
			}

			next.ServeHTTP(w, req)
		},
	)
}

func withTimeout(duration time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			ctx, cancel := context.WithTimeout(req.Context(), duration)
			defer cancel()

			next.ServeHTTP(w, req.WithContext(ctx))
		},
	)
}

func main() {
	log.Println("Starting the server...")

	itemRepository, err := postgres.NewItemRepository(postgresURL)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle(
		"/item",
		post(
			withTimeout(
				200*time.Millisecond,
				item.NewCreateController(
					createItem.NewHandler(
						createItem.Validator,
						itemRepository,
					),
				),
			),
		),
	)
	http.ListenAndServe(":8080", nil)
}
