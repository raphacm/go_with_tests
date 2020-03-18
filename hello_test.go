package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("should return Hello to a person", func(t *testing.T) {
		got := Hello("Raphael")
		want := "Hello Raphael"

		assertMessage(t, got, want)
	})

	t.Run("should return Hello World when a empty string is passed as parameter", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertMessage(t, got, want)
	})
}
