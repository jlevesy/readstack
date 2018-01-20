package errors

import (
	"net/http"
)

// HandlerStub is a stub enabling to define behaviour of
// errors.Handler on testsuite
type HandlerStub struct {
	OnHandle func(w http.ResponseWriter, err error)
}

func (h *HandlerStub) Handle(w http.ResponseWriter, err error) {
	h.OnHandle(w, err)
}
