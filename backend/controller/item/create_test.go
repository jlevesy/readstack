package item

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jlevesy/readstack/handler/item/create"
	errorsStub "github.com/jlevesy/readstack/test/stub/controller/errors"
	handlerStub "github.com/jlevesy/readstack/test/stub/handler/item/create"
)

func TestItShouldHandleErrorOnDecodeError(t *testing.T) {
	input := bytes.NewBufferString("bozoleclown")
	req := httptest.NewRequest("POST", "/api/v1/item", input)
	w := httptest.NewRecorder()

	errorCalledCount := 0
	errorHandler := errorsStub.HandlerStub{
		OnHandle: func(w http.ResponseWriter, err error) {
			errorCalledCount++
		},
	}

	handleCalledCount := 0
	handler := handlerStub.HandlerStub{
		OnHandle: func(ctx context.Context, req *create.Request) error {
			handleCalledCount++
			return nil
		},
	}

	subject := NewCreateController(&handler, &errorHandler)

	subject.ServeHTTP(w, req)

	if errorCalledCount != 1 {
		t.Fatalf("Expected 1 call to error handler, got %d", errorCalledCount)
	}

	if handleCalledCount != 0 {
		t.Fatalf("Expected no call to handler, got %d", handleCalledCount)
	}
}

func TestItShouldHandleHandlerError(t *testing.T) {
	input := bytes.NewBuffer([]byte{})
	json.NewEncoder(input).Encode(&create.Request{
		Name: "John",
		URL:  "Doe",
	})

	req := httptest.NewRequest("POST", "/api/v1/item", input)
	w := httptest.NewRecorder()

	errorCalledCount := 0
	errorHandler := errorsStub.HandlerStub{
		OnHandle: func(w http.ResponseWriter, err error) {
			errorCalledCount++
		},
	}

	handleCalledCount := 0
	handler := handlerStub.HandlerStub{
		OnHandle: func(ctx context.Context, req *create.Request) error {
			handleCalledCount++
			return errors.New("(╯°o°）╯ ┻━┻")
		},
	}

	subject := NewCreateController(&handler, &errorHandler)

	subject.ServeHTTP(w, req)

	if errorCalledCount != 1 {
		t.Fatalf("Expected 1 call to error handler, got %d", errorCalledCount)
	}

	if handleCalledCount != 1 {
		t.Fatalf("Expected one call to handler, got %d", handleCalledCount)
	}
}

func TestItShouldSetStatusCreatedOnSuccess(t *testing.T) {
	input := bytes.NewBuffer([]byte{})
	json.NewEncoder(input).Encode(&create.Request{
		Name: "John",
		URL:  "Doe",
	})

	req := httptest.NewRequest("POST", "/api/v1/item", input)
	w := httptest.NewRecorder()

	errorCalledCount := 0
	errorHandler := errorsStub.HandlerStub{
		OnHandle: func(w http.ResponseWriter, err error) {
			errorCalledCount++
		},
	}

	handleCalledCount := 0
	handler := handlerStub.HandlerStub{
		OnHandle: func(ctx context.Context, req *create.Request) error {
			handleCalledCount++
			return nil
		},
	}

	subject := NewCreateController(&handler, &errorHandler)

	subject.ServeHTTP(w, req)

	if errorCalledCount != 0 {
		t.Fatalf("Expected no call to error handler, got %d", errorCalledCount)
	}

	if handleCalledCount != 1 {
		t.Fatalf("Expected one call to handler, got %d", handleCalledCount)
	}

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected HTTP status created, got %d", w.Code)
	}
}
