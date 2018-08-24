package integration

import (
	"context"
	"reflect"
	"testing"

	"github.com/jlevesy/readstack/server/item"
)

func setup() (item.Repository, func(), error) {
	db, done, err := SetupDB()

	if err != nil {
		return nil, nil, err
	}

	repo, err := item.NewSQLRepository(db)

	if err != nil {
		return nil, nil, err
	}

	return repo, done, nil
}

func TestRepository(t *testing.T) {
	repo, done, err := setup()

	if err != nil {
		t.Fatal(err)
	}

	defer done()

	originalItem := item.New("foo", "https://foo.bar.com")

	if err := repo.Create(context.Background(), originalItem); err != nil {
		t.Fatal(err)
	}

	if originalItem.GetID() != 1 {
		t.Errorf("Expected ID 1 got %d", originalItem.GetID())
	}

	items, err := repo.FindAll(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if len(items) != 1 {
		t.Errorf("Invalid count of items, expected 1 got %d", len(items))
	}

	item := items[0]

	if !reflect.DeepEqual(item, originalItem) {
		t.Errorf("Read item is not deeply equal to written item")
	}

	err = repo.Delete(context.Background(), item)

	if err != nil {
		t.Fatal(err)
	}

	items, err = repo.FindAll(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if len(items) != 0 {
		t.Errorf("Invalid count of items, expected 0 items got  %d", len(items))
	}
}
