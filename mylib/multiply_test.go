package mylib

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	reference := 5 * 4
	result := Multiply(5, 4)

	if result != reference {
		t.Errorf("Multiply should return %d instead of %d", reference, result)
	}
}
