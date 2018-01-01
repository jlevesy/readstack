package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/handler/item/create"
)

type createController struct {
	handler create.Handler
}

func NewCreateController(handler create.Handler) http.Handler {
	return &createController{handler}
}

func (c *createController) ServeHTTP(w http.ResponseWriter, httpReq *http.Request) {
	defer httpReq.Body.Close()

	var req create.Request

	if err := json.NewDecoder(httpReq.Body).Decode(&req); err != nil {
		// TODO factorize errors
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := c.handler.Handle(httpReq.Context(), &req); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
