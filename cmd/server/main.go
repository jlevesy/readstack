package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository/postgres"
)

const (
	postgresURL = "postgres://readstack:notsecret@db:5432/readstack?sslmode=disable"
)

func createItem() (http.Handler, error) {
	repository, err := postgres.NewItemRepository(postgresURL)

	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		defer r.Body.Close()

		i := model.Item{}

		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}

		if err := repository.Save(ctx, &i); err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
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
