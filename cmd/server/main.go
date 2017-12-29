package main

import (
	"log"
	"net/http"

	// "github.com/jlevesy/readstack/model"
	"github.com/jlevesy/readstack/repository/postgres"
)

const (
	postgresUrl = "postgres://readstack:notsecret@db:5432/readstack?sslmode=disable"
)

func createArticle() (http.Handler, error) {
	_, err := postgres.NewArticleRepository(postgresUrl)

	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}), nil
}

func main() {
	log.Println("Starting the server...")

	createArticle, err := createArticle()

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/article", createArticle)
}
