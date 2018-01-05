package item

import (
	"net/http"
)

type indexController struct{}

func NewIndexController() http.Handler {
	return &indexController{}
}

func (i *indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
