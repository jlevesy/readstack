package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller/errors"
	"github.com/jlevesy/readstack/handler/item/index"
)

type indexController struct {
	handler    index.Handler
	errHandler errors.HttpErrorHandler
}

func NewIndexController(handler index.Handler, errHandler errors.HttpErrorHandler) http.Handler {
	return &indexController{handler, errHandler}
}

func (i *indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := i.handler.Handle(r.Context())

	if err != nil {
		i.errHandler.HandleHttpError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		i.errHandler.HandleHttpError(w, err)
		return
	}
}
