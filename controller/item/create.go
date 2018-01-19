package item

import (
	"encoding/json"
	"net/http"

	"github.com/jlevesy/readstack/controller/errors"
	"github.com/jlevesy/readstack/handler/item/create"
)

type createController struct {
	handler    create.Handler
	errHandler errors.HttpErrorHandler
}

func NewCreateController(handler create.Handler, errHandler errors.HttpErrorHandler) http.Handler {
	return &createController{handler, errHandler}
}

func (c *createController) ServeHTTP(w http.ResponseWriter, httpReq *http.Request) {
	defer httpReq.Body.Close()

	var req create.Request

	if err := json.NewDecoder(httpReq.Body).Decode(&req); err != nil {
		c.errHandler.HandleHttpError(w, err)
		return
	}

	if err := c.handler.Handle(httpReq.Context(), &req); err != nil {
		c.errHandler.HandleHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
