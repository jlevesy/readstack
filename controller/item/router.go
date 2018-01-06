package item

import (
	"github.com/gorilla/mux"

	"github.com/jlevesy/readstack/repository"

	createItem "github.com/jlevesy/readstack/handler/item/create"
	indexItem "github.com/jlevesy/readstack/handler/item/index"
)

func MountRoutes(router *mux.Router, itemRepository repository.ItemRepository) {
	router.Path("/item").Methods("POST").Handler(
		NewCreateController(
			createItem.NewHandler(
				createItem.Validator,
				itemRepository,
			),
		),
	)

	router.Path("/item").Methods("GET").Handler(
		NewIndexController(
			indexItem.NewHandler(itemRepository),
		),
	)
}
