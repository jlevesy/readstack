package index

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/jlevesy/readstack/model"

	"github.com/jlevesy/readstack/test/stub/repository"
)

func TestItShouldReturnAllResults(t *testing.T) {
	dbResults := []*model.Item{
		model.NewItem("foo", "bar"),
		model.NewItem("foo", "bar"),
		model.NewItem("foo", "bar"),
	}

	mockRepository := &repository.ItemRepositoryStub{
		OnFindAll: func(context.Context) ([]*model.Item, error) {
			return dbResults, nil
		},
	}

	subject := NewHandler(mockRepository)

	res, err := subject.Handle(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if !reflect.DeepEqual(res.Items, dbResults) {
		t.Fatalf("Expected returned results to be deep equals")
	}
}

func TestItForwardsARepositoryError(t *testing.T) {
	returnedErr := errors.New("Failed to reach database.")

	mockRepository := &repository.ItemRepositoryStub{
		OnFindAll: func(context.Context) ([]*model.Item, error) {
			return []*model.Item{}, returnedErr
		},
	}

	subject := NewHandler(mockRepository)

	res, err := subject.Handle(context.Background())

	if err == nil {
		t.Fatal("Expected an error, got nothng")
	}

	if res != nil {
		t.Fatalf("Expected no response, got %v", res)
	}

	if err != returnedErr {
		t.Fatalf("Expected an error %s, got %s", err, returnedErr)
	}
}
