package validation

import (
	"testing"
)

func TestRequireHTTPURL(t *testing.T) {
	cases := []struct {
		Label            string
		FieldName        string
		Value            string
		ShouldRaiseError bool
	}{
		{
			"ShouldRaiseErrorOnEmptyValue",
			"Dummy",
			"",
			true,
		},
		{
			"SouldRaiseErrorIfValueIsBlank",
			"Dummy",
			"     ",
			true,
		},
		{
			"SouldRaiseErrorIfURLIsNotHTTP",
			"Dummy",
			"postgres://foo.bar.com/dbtoto",
			true,
		},
		{
			"ShouldNotRaiseErrorIfValueIsHTTPURL",
			"Dummy",
			"https://foo.bar.com/dbtoto",
			false,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Label, func(t *testing.T) {
			res := RequireHTTPURL(testCase.FieldName, testCase.Value)

			if !testCase.ShouldRaiseError && res != nil {
				t.Fatalf("Expected no error, got %v", res)
			}

			if testCase.ShouldRaiseError && res == nil {
				t.Fatal("Expected an error, got nothing")
			}
		})
	}
}
