package create

import (
	"context"
	"errors"
	"testing"

	rsError "github.com/jlevesy/readstack/error"
	"github.com/jlevesy/readstack/model"

	"github.com/jlevesy/readstack/test/stub/repository"
)

func TestItCreatesAndSavesAnItem(t *testing.T) {
	request := NewRequest("Name", "https://name.com")
	var savedItem *model.Item

	mockRepository := repository.ItemRepositoryStub{
		OnSave: func(ctx context.Context, i *model.Item) error {
			savedItem = i
			return nil
		},
	}

	validator := func(r *Request) []*rsError.Violation {
		return []*rsError.Violation{}
	}

	subject := NewHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), request)

	if err != nil {
		t.Fatalf("Expected no errors, got %v", err)
	}

	if savedItem.URL != request.URL {
		t.Fatalf("Invalid saved URL, expectd %s got %s", request.URL, savedItem.URL)
	}

	if savedItem.Name != request.Name {
		t.Fatalf("Invalid saved Name, expectd %s got %s", request.Name, savedItem.Name)
	}
}

func TestItReportsAValdationError(t *testing.T) {
	request := NewRequest("Name", "https://name.com")
	var savedItem *model.Item

	mockRepository := repository.ItemRepositoryStub{
		OnSave: func(ctx context.Context, i *model.Item) error {
			savedItem = i
			return nil
		},
	}

	validator := func(r *Request) []*rsError.Violation {
		return []*rsError.Violation{
			{
				Name:   "Foo",
				Reason: "Bar",
			},
		}
	}

	subject := NewHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), request)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if _, ok := err.(*rsError.ValidationError); !ok {
		t.Fatalf("Expected a validator error, got %T", err)
	}
}

func TestItReportsARepositoryError(t *testing.T) {
	request := NewRequest("Name", "https://name.com")
	returnedErr := errors.New("Failed to lalala the database")

	mockRepository := repository.ItemRepositoryStub{
		OnSave: func(ctx context.Context, i *model.Item) error {
			return returnedErr
		},
	}

	validator := func(r *Request) []*rsError.Violation {
		return []*rsError.Violation{}
	}

	subject := NewHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), request)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if returnedErr != err {
		t.Fatalf("Expectedr error %s, got %s", returnedErr, err)
	}
}
