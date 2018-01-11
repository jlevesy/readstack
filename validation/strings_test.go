package validation

import (
	"testing"
)

func TestRequireNotBlank(t *testing.T) {
	cases := []struct {
		Label            string
		FieldName        string
		Value            string
		ShouldRaiseError bool
	}{
		{
			"ShouldNotRaiseErrorIfValueIsNotBlank",
			"Dummy",
			"foo",
			false,
		},
		{
			"SouldRaiseErrorIfValueIsBlank",
			"Dummy",
			"     ",
			true,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Label, func(t *testing.T) {
			res := RequireNotBlank(testCase.FieldName, testCase.Value)

			if !testCase.ShouldRaiseError && res != nil {
				t.Fatalf("Expected no error, got %v", res)
			}

			if testCase.ShouldRaiseError && res == nil {
				t.Fatal("Expected an error, got nothing")
			}
		})
	}
}
