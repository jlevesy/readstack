package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jlevesy/readstack/handler/errors"
	"github.com/jlevesy/readstack/logger"
)

const (
	validationErrorType     = "urn:api:readstack:error:validation-error"
	jsonErrorType           = "urn:api:readstack:error:json-error"
	internalServerErrorType = "urn:api:readstack:error:internal-server-error"

	errorDefaultMessage = "Something went terribly wrong"
)

type invalidParam struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type apiError struct {
	Type          string          `json:"type"`
	Title         string          `json:"title"`
	InvalidParams []*invalidParam `json:"invalid-params,omitempty"`
}

// HttpErrorHandler defines an interface to handle errors at controller level
type HttpErrorHandler interface {
	HandleHttpError(w http.ResponseWriter, err error)
}

type httpErrorHandler struct {
	logger logger.Logger
}

// NewHttpErrorHandler returns an HttpErrorHandler
func NewHttpErrorHandler(logger logger.Logger) HttpErrorHandler {
	return &httpErrorHandler{logger}
}

func (h *httpErrorHandler) HandleHttpError(w http.ResponseWriter, err error) {
	h.logger.Error("Handler error: [%T] %s", err, err.Error())

	switch v := err.(type) {
	case *json.SyntaxError:
		handleError(
			w,
			&apiError{
				Type:  jsonErrorType,
				Title: fmt.Sprintf("Failed to decode request : %s", err),
			},
			http.StatusUnprocessableEntity,
		)
	case *errors.ValidationError:
		handleValidationError(w, v, http.StatusBadRequest)
	default:
		handleError(
			w,
			&apiError{
				Type:  internalServerErrorType,
				Title: errorDefaultMessage,
			},
			http.StatusInternalServerError,
		)
	}
}

func handleValidationError(w http.ResponseWriter, err *errors.ValidationError, statusCode int) {
	invalidParams := make([]*invalidParam, len(err.Violations))

	for i, v := range err.Violations {
		invalidParams[i] = &invalidParam{
			Name:   v.Name,
			Reason: v.Reason,
		}
	}

	handleError(
		w,
		&apiError{
			Type:          validationErrorType,
			Title:         "Failed to validate request",
			InvalidParams: invalidParams,
		},
		statusCode,
	)
}

func handleError(w http.ResponseWriter, err *apiError, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}
