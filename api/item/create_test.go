package item

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/jlevesy/readstack/api/validation"
)

func TestValidator(t *testing.T) {
	cases := []struct {
		Input        *CreateRequest
		Expectations []*validation.Violation
	}{
		{
			Input:        &CreateRequest{"Foo", "https://foo.bar.com"},
			Expectations: []*validation.Violation{},
		},
		{
			Input: &CreateRequest{"", "https://foo.bar.com"},
			Expectations: []*validation.Violation{
				{
					Name:   "Name",
					Reason: "Should not be blank",
				},
			},
		},
		{
			Input: &CreateRequest{"Bar", ""},
			Expectations: []*validation.Violation{
				{
					Name:   "URL",
					Reason: "Should not be blank",
				},
				{
					Name:   "URL",
					Reason: "Unsuported URL scheme, only http and https are allowed",
				},
			},
		},
	}

	for _, testCase := range cases {
		t.Run("", func(t *testing.T) {
			violations := CreateValidator(testCase.Input)

			for i, v := range violations {
				expectation := testCase.Expectations[i]
				if !reflect.DeepEqual(*v, *expectation) {
					t.Errorf("Expected %v, got %v", *expectation, *v)
				}
			}
		})
	}
}

func TestItCreatesAndSavesAnItem(t *testing.T) {
	request := CreateRequest{"Name", "https://name.com"}
	var createdItem *Model

	mockRepository := RepositoryStub{
		OnCreate: func(ctx context.Context, i *Model) error {
			createdItem = i
			return nil
		},
	}

	validator := func(r *CreateRequest) []*validation.Violation {
		return []*validation.Violation{}
	}

	subject := NewCreateHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), &request)

	if err != nil {
		t.Fatalf("Expected no errors, got %v", err)
	}

	if createdItem.URL != request.URL {
		t.Fatalf("Invalid saved URL, expectd %s got %s", request.URL, createdItem.URL)
	}

	if createdItem.Name != request.Name {
		t.Fatalf("Invalid saved Name, expectd %s got %s", request.Name, createdItem.Name)
	}
}

func TestItReportsAValdationError(t *testing.T) {
	request := CreateRequest{"Name", "https://name.com"}

	mockRepository := RepositoryStub{
		OnCreate: func(ctx context.Context, i *Model) error {
			return nil
		},
	}

	validator := func(r *CreateRequest) []*validation.Violation {
		return []*validation.Violation{
			{
				Name:   "Foo",
				Reason: "Bar",
			},
		}
	}

	subject := NewCreateHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), &request)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if _, ok := err.(*validation.Error); !ok {
		t.Fatalf("Expected a validator error, got %T", err)
	}
}

func TestItReportsARepositoryError(t *testing.T) {
	request := CreateRequest{"Name", "https://name.com"}
	returnedErr := errors.New("Failed to lalala the database")

	mockRepository := RepositoryStub{
		OnCreate: func(ctx context.Context, i *Model) error {
			return returnedErr
		},
	}

	validator := func(r *CreateRequest) []*validation.Violation {
		return []*validation.Violation{}
	}

	subject := NewCreateHandler(validator, &mockRepository)

	err := subject.Handle(context.Background(), &request)

	if err == nil {
		t.Fatal("Expected an error, got nothing")
	}

	if returnedErr != err {
		t.Fatalf("Expectedr error %s, got %s", returnedErr, err)
	}
}
