package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	readstackError "github.com/jlevesy/readstack/error"
)

func HandleError(w http.ResponseWriter, err error) {
	log.Printf("Handler error: [%T] %s", err, err.Error())

	switch v := err.(type) {
	case *json.SyntaxError:
		http.Error(
			w,
			fmt.Sprintf("Failed to decode request %v", v),
			http.StatusUnprocessableEntity,
		)
	case *readstackError.ValidationError:
		http.Error(
			w,
			fmt.Sprintf("Failed to validate request %v", err),
			http.StatusBadRequest,
		)
	default:
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
