package item

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jlevesy/readstack/handler/item/index"
	"github.com/jlevesy/readstack/model"

	errorsStub "github.com/jlevesy/readstack/test/stub/controller/errors"
	handlerStub "github.com/jlevesy/readstack/test/stub/handler/item/index"
)

func TestItShouldHandleIndexHandlerError(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/item", nil)
	w := httptest.NewRecorder()

	errorCalledCount := 0
	errorHandler := errorsStub.HTTPErrorHandlerStub{
		OnHandleHTTPError: func(w http.ResponseWriter, err error) {
			errorCalledCount++
		},
	}

	handleCalledCount := 0
	handler := handlerStub.HandlerStub{
		OnHandle: func(ctx context.Context) (*index.Response, error) {
			handleCalledCount++
			return nil, errors.New("¯\\_(ツ)_/¯")
		},
	}

	subject := NewIndexController(&handler, &errorHandler)

	subject.ServeHTTP(w, req)

	if errorCalledCount != 1 {
		t.Fatalf("Expected 1 call to error handler, got %d", errorCalledCount)
	}

	if handleCalledCount != 1 {
		t.Fatalf("Expected no call to handler, got %d", handleCalledCount)
	}
}

func TestItShouldReturnIndexHandlerResponse(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/item", nil)
	w := httptest.NewRecorder()

	errorCalledCount := 0
	errorHandler := errorsStub.HTTPErrorHandlerStub{
		OnHandleHTTPError: func(w http.ResponseWriter, err error) {
			errorCalledCount++
		},
	}

	expectedResponse := index.Response{
		Items: []*model.Item{
			{
				Name: "Foo",
				URL:  "Bar",
			},
		},
	}
	handleCalledCount := 0
	handler := handlerStub.HandlerStub{
		OnHandle: func(ctx context.Context) (*index.Response, error) {
			handleCalledCount++
			return &expectedResponse, nil
		},
	}

	subject := NewIndexController(&handler, &errorHandler)

	subject.ServeHTTP(w, req)

	if errorCalledCount != 0 {
		t.Fatalf("Expected no call to error handler, got %d", errorCalledCount)
	}

	if handleCalledCount != 1 {
		t.Fatalf("Expected a call to handler, got %d", handleCalledCount)
	}

	if w.Code != http.StatusOK {
		t.Fatalf("Expected a statusOK, got %v", w.Code)
	}

	response := index.Response{}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Expected returned payload to be JSON, failed to decode it %v", err)
	}

	for i, expectedItem := range expectedResponse.Items {
		item := response.Items[i]

		if !reflect.DeepEqual(expectedItem, item) {
			t.Fatalf("Failed to compare item, expected %v got %v", expectedItem, item)
		}
	}
}
