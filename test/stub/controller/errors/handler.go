package errors

import (
	"net/http"
)

// HTTPErrorHandlerStub is a stub enabling to define behaviour of
// errors.HttpErrorHandler on testsuite
type HTTPErrorHandlerStub struct {
	OnHandleHTTPError func(w http.ResponseWriter, err error)
}

func (h *HTTPErrorHandlerStub) HandleHTTPError(w http.ResponseWriter, err error) {
	h.OnHandleHTTPError(w, err)
}
