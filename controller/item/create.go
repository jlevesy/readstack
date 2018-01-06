package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller"
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
		controller.HandleError(w, err)
		return
	}

	if err := c.handler.Handle(httpReq.Context(), &req); err != nil {
		controller.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
