package errors

import (
	"testing"
)

func TestItCanBeSerializedAsAString(t *testing.T) {

	v := NewValidationError(
		[]*Violation{
			{Name: "Foo", Reason: "Bar"},
			{Name: "Foo", Reason: "Bar"},
		},
	)

	res := v.Error()

	if res != "Foo : Bar, Foo : Bar" {
		t.Fatalf("Expected errors to be serialized")
	}

}
