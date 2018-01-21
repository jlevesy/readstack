package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller/errors"
	"github.com/jlevesy/readstack/handler/item/index"
)

type indexController struct {
	handler    index.Handler
	errHandler errors.Handler
}

// NewIndexController returns an instance of an indexController as an http.Handler
func NewIndexController(handler index.Handler, errHandler errors.Handler) http.Handler {
	return &indexController{handler, errHandler}
}

func (i *indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := i.handler.Handle(r.Context())

	if err != nil {
		i.errHandler.Handle(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
