package assert

import (
	"reflect"
	"testing"
)

var t *testing.T

func DeepEqual[T any | ~[]any](actual T, expected T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Found %v, expected %v", actual, expected)
	}
}

func Equal[T comparable](actual T, expected T) {
	if actual != expected {
		t.Fatalf("%v does not equal %v", actual, expected)
	}
}
