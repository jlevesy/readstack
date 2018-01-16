package create

import (
	"reflect"
	"testing"

	"github.com/jlevesy/readstack/error"
)

func TestValidator(t *testing.T) {
	cases := []struct {
		Input        *Request
		Expectations []*error.Violation
	}{
		{
			Input:        NewRequest("Foo", "https://foo.bar.com"),
			Expectations: []*error.Violation{},
		},
		{
			Input: NewRequest("", "https://foo.bar.com"),
			Expectations: []*error.Violation{
				{
					Name:   "Name",
					Reason: "Should not be blank",
				},
			},
		},
		{
			Input: NewRequest("Bar", ""),
			Expectations: []*error.Violation{
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
			violations := Validator(testCase.Input)

			for i, v := range violations {
				expectation := testCase.Expectations[i]
				if !reflect.DeepEqual(*v, *expectation) {
					t.Errorf("Expected %v, got %v", *expectation, *v)
				}
			}
		})
	}
}
