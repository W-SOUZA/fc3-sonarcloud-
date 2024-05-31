// package main

// import "testing"

// func TestSum(t *testing.T) {
// 	result := sum(2, 3)
// 	if result != 5 {
// 		t.Error("The result must be 5")
// 	}
// }

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add function returned incorrect result, got: %d, want: %d.", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract function returned incorrect result, got: %d, want: %d.", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := multiply(3, 4)
	expected := 12
	if result != expected {
		t.Errorf("Multiply function returned incorrect result, got: %d, want: %d.", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := divide(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("Divide function returned incorrect result, got: %d, want: %d.", result, expected)
	}
}

func TestModulo(t *testing.T) {
	result := modulo(10, 3)
	expected := 1
	if result != expected {
		t.Errorf("Modulo function returned incorrect result, got: %d, want: %d.", result, expected)
	}
}
