package main

import "testing"

func TestDivide(t *testing.T) {
	result := divide(1, 1)
	if result != 1 {
		t.Error("Expected 1/1 = 1")
	}
	result = divide(1, 2)
	if result != 0 {
		t.Error("Expected 1/2 = 0")
	}
	result = divide(2, 0)
	if result != 2 {
		t.Error("Expected 1/2 = 0")
	}

}
