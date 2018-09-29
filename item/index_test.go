package item

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestItShouldReturnAllResults(t *testing.T) {
	dbResults := []*Model{
		{0, "foo", "bar"},
		{1, "foo", "bar"},
		{2, "foo", "bar"},
	}

	mockRepository := &RepositoryStub{
		OnFindAll: func(context.Context) ([]*Model, error) {
			return dbResults, nil
		},
	}

	subject := NewIndexHandler(mockRepository)

	res, err := subject.Handle(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if !reflect.DeepEqual(res.Items, dbResults) {
		t.Fatalf("Expected returned results to be deep equals")
	}
}

func TestItForwardsARepositoryError(t *testing.T) {
	returnedErr := errors.New("failed to reach database")

	mockRepository := &RepositoryStub{
		OnFindAll: func(context.Context) ([]*Model, error) {
			return []*Model{}, returnedErr
		},
	}

	subject := NewIndexHandler(mockRepository)

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
