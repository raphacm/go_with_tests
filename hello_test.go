package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Raphael")
	want := "Hello Raphael"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
