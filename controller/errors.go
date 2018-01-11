package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	readstackError "github.com/jlevesy/readstack/error"
)

const (
	validationErrorType     = "urn:api:readstack:error:validation-error"
	jsonErrorType           = "urn:api:readstack:error:json-error"
	internalServerErrorType = "urn:api:readstack:error:internal-server-error"
)

type invalidParam struct {
	Name   string
	Reason string
}

type apiError struct {
	Type          string
	Title         string
	InvalidParams []*invalidParam
}

func HandleError(w http.ResponseWriter, err error) {
	log.Printf("Handler error: [%T] %s", err, err.Error())

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
	case *readstackError.ValidationError:
		handleValidationError(w, v, http.StatusBadRequest)
	default:
		handleError(
			w,
			&apiError{
				Type:  internalServerErrorType,
				Title: "Something went terribly wrong.",
			},
			http.StatusInternalServerError,
		)
	}
}

func handleValidationError(w http.ResponseWriter, err *readstackError.ValidationError, statusCode int) {
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
	json.NewEncoder(w).Encode(err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
}
