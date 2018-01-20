package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller/errors"
	"github.com/jlevesy/readstack/handler/item/index"
)

type indexController struct {
	handler    index.Handler
	errHandler errors.HTTPErrorHandler
}

// NewIndexController returns an instance of an indexController as an http.Handler
func NewIndexController(handler index.Handler, errHandler errors.HTTPErrorHandler) http.Handler {
	return &indexController{handler, errHandler}
}

func (i *indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := i.handler.Handle(r.Context())

	if err != nil {
		i.errHandler.HandleHTTPError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)

	w.WriteHeader(http.StatusOK)
}
