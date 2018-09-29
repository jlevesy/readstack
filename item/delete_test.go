package item

import (
	"context"
	"errors"
	"testing"
)

func TestItDeletesAnItem(t *testing.T) {
	request := DeleteRequest{10}
	var deletedItem *Model

	mockRepository := RepositoryStub{
		OnDelete: func(ctx context.Context, i *Model) error {
			deletedItem = i
			return nil
		},
	}

	subject := NewDeleteHandler(&mockRepository)

	err := subject.Handle(context.Background(), &request)

	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}

	if request.ID != deletedItem.GetID() {
		t.Fatalf("Wrong ID forwared to the repository")
	}
}

func TestItForwardsAnError(t *testing.T) {
	request := DeleteRequest{10}

	mockRepository := RepositoryStub{
		OnDelete: func(ctx context.Context, i *Model) error {
			return errors.New("Nope")
		},
	}

	subject := NewDeleteHandler(&mockRepository)

	err := subject.Handle(context.Background(), &request)

	if err == nil {
		t.Fatalf("Expected an error, got nothing")
	}

}
