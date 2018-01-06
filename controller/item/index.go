package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller"
	"github.com/jlevesy/readstack/handler/item/index"
)

type indexController struct {
	handler index.Handler
}

func NewIndexController(handler index.Handler) http.Handler {
	return &indexController{handler}
}

func (i *indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := i.handler.Handle(r.Context())

	if err != nil {
		controller.HandleError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		controller.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
