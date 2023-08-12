package main

import "testing"

func TestIsTrue(t *testing.T) {
	if !isTrue(true) {
		t.Error("Expected true, got false")
	}
}