package main

import (
	"log"
	"net/http"

	// "github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository/postgres"
)

const (
	postgresURL = "postgres://readstack:notsecret@db:5432/readstack?sslmode=disable"
)

func createItem() (http.Handler, error) {
	_, err := postgres.NewItemRepository(postgresURL)

	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}), nil
}

func main() {
	log.Println("Starting the server...")

	createItem, err := createItem()

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/item", createItem)
	http.ListenAndServe(":8080", nil)
}
