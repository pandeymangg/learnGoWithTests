package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Anshuman")

	want := "Hello, Anshuman!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
