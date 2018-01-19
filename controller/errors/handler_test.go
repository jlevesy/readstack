package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	handlerErrors "github.com/jlevesy/readstack/handler/errors"
	"github.com/jlevesy/readstack/test/stub/logger"
)

func TestItCanHandleErrors(t *testing.T) {
	testCases := []struct {
		Label              string
		Err                error
		ExpectedStatusCode int
		Expectation        apiError
	}{
		{
			"WithRandomError",
			errors.New("YOLO"),
			http.StatusInternalServerError,
			apiError{
				Type:  internalServerErrorType,
				Title: errorDefaultMessage,
			},
		},
		{
			"WithJSONError",
			&json.SyntaxError{},
			http.StatusUnprocessableEntity,
			apiError{
				Type:  jsonErrorType,
				Title: "Failed to decode request : ",
			},
		},
		{
			"WithValidationError",
			&handlerErrors.ValidationError{
				[]*handlerErrors.Violation{
					{
						Name:   "Foo",
						Reason: "Bar",
					},
				},
			},
			http.StatusBadRequest,
			apiError{
				Type:  validationErrorType,
				Title: "Failed to validate request",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Label, func(t *testing.T) {
			loggerCalledCount := 0
			logger := logger.LoggerStub{
				OnError: func(format string, args ...interface{}) {
					loggerCalledCount++
				},
			}
			handler := NewHttpErrorHandler(&logger)
			resWriter := httptest.NewRecorder()

			handler.HandleHttpError(resWriter, testCase.Err)

			if loggerCalledCount != 1 {
				t.Fatalf("Expected 1 call to logger, got %d", loggerCalledCount)
			}

			if testCase.ExpectedStatusCode != resWriter.Code {
				t.Fatalf("Expected status code %d, got %d", testCase.ExpectedStatusCode, resWriter.Code)
			}

			parsed := apiError{}
			err := json.NewDecoder(resWriter.Body).Decode(&parsed)

			if err != nil {
				t.Fatalf("Failed to decode JSON output, got %v", err)
			}

			if parsed.Type != testCase.Expectation.Type {
				t.Fatalf(
					"Invalid error type, expected %s, got %s",
					testCase.Expectation.Type,
					parsed.Type,
				)
			}

			if parsed.Title != testCase.Expectation.Title {
				t.Fatalf(
					"Invalid title, expected %s, got %s",
					testCase.Expectation.Title,
					parsed.Title,
				)
			}

			for i, expectedInvalidParams := range testCase.Expectation.InvalidParams {
				resInvalidParams := parsed.InvalidParams[i]

				if !reflect.DeepEqual(expectedInvalidParams, resInvalidParams) {
					t.Fatalf(
						"Unexpected invalid param, expected %v, got %v",
						expectedInvalidParams,
						resInvalidParams,
					)
				}
			}
		})
	}
}
