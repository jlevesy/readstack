package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlevesy/envconfig"

	"github.com/jlevesy/readstack/middleware"
	"github.com/jlevesy/readstack/repository"
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

func router(itemRepository repository.ItemRepository) http.Handler {
	r := mux.NewRouter()

	r.Path("/item").Methods("POST").Handler(
		item.NewCreateController(
			createItem.NewHandler(
				createItem.Validator,
				itemRepository,
			),
		),
	)

	return r
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

	http.ListenAndServe(
		":8080",
		middleware.WithInMemoryTimingRecorder(
			middleware.Timeout(
				200*time.Millisecond,
				middleware.RequestLogger(
					middleware.RecordDuration(
						middleware.HandlerDuration,
						router(itemRepository),
					),
				),
			),
		),
	)
}
