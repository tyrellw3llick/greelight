package internal

import (
	"testing"
)

func Assert[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
