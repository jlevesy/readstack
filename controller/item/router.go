package item

import (
	"github.com/gorilla/mux"

	"github.com/jlevesy/readstack/controller/errors"
	createItem "github.com/jlevesy/readstack/handler/item/create"
	indexItem "github.com/jlevesy/readstack/handler/item/index"
	"github.com/jlevesy/readstack/repository"
)

// MountRoutes map item related controllers to given router
func MountRoutes(router *mux.Router, itemRepository repository.ItemRepository, errorHandler errors.Handler) {
	router.Path("/item").Methods("POST").Handler(
		NewCreateController(
			createItem.NewHandler(
				createItem.Validator,
				itemRepository,
			),
			errorHandler,
		),
	)

	router.Path("/item").Methods("GET").Handler(
		NewIndexController(
			indexItem.NewHandler(itemRepository),
			errorHandler,
		),
	)
}
