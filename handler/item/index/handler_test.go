package index

import (
	"context"
	"errors"
	"testing"

	"github.com/jlevesy/readstack/model"

	"github.com/jlevesy/readstack/test/stub/repository"
)

func TestIndexHandler(t *testing.T) {
	cases := []struct {
		Name           string
		OnFindAll      func(context.Context) ([]*model.Item, error)
		ShouldFail     bool
		ExpectedLength int
	}{
		{
			"itShouldReturnResults",
			func(context.Context) ([]*model.Item, error) {
				return []*model.Item{
					model.NewItem("foo", "bar"),
					model.NewItem("foo", "bar"),
					model.NewItem("foo", "bar"),
				}, nil
			},
			false,
			3,
		},
		{
			"itShouldForwardFailure",
			func(context.Context) ([]*model.Item, error) {
				return []*model.Item{}, errors.New("YOLO")
			},
			true,
			0,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			mockRepository := repository.ItemRepositoryStub{
				OnFindAll: testCase.OnFindAll,
			}

			subject := NewHandler(&mockRepository)

			res, err := subject.Handle(context.Background())

			if testCase.ShouldFail {
				if err == nil {
					t.Fatal("Expected an error, got nothing")
				}

				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if len(res.Items) != testCase.ExpectedLength {
				t.Fatalf("Expected %d, got %d", testCase.ExpectedLength, len(res.Items))
			}

		})
	}

}
