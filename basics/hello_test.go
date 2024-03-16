package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gupta")
	want := "Hello Gupta"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
