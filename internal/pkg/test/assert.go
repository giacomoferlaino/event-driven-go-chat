package test

import (
	"testing"
)

func AssertEqual[T comparable](want T, got T, t *testing.T) {
	if want != got {
		t.Errorf("Assertion error, want: %v, got: %v", want, got)
	}
}

func AssertNotEqual[T comparable](want T, got T, t *testing.T) {
	if want == got {
		t.Errorf("Assertion error, the two parameters should be different.\nReceived value: %v", got)
	}
}
