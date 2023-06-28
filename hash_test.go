package main

import "testing"

func TestHash(t *testing.T) {
	gotHash := hash("hello")
	if gotHash != 42 {
		t.Errorf("got %v, want 42", gotHash)
	}
}
