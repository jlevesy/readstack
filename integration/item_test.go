package integration

import (
	"context"
	"testing"

	"github.com/jlevesy/readstack/api"
)

func TestIntexReturnsAllTheItems(t *testing.T) {
	tc := setup(t)
	defer tc.TearDown()

	res, err := tc.ItemClient.Index(context.Background(), &api.IndexRequest{})

	if err != nil {
		t.Fatal(err)
	}

	if len(res.Items) != 0 {
		t.Error("Expected no items")
	}
}
