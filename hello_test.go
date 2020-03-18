package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("should return Hello to a person", func(t *testing.T) {
		got := Hello("Raphael")
		want := "Hello Raphael"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

	t.Run("should return Hello World when a empty string is passed as parameter", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})
}
