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
		got := Hello("Raphael", "English")
		want := "Hello Raphael"

		assertMessage(t, got, want)
	})

	t.Run("should return Hello World when a empty string is passed as name parameter", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"

		assertMessage(t, got, want)
	})

	t.Run("should return greeting in Spanish when a language parameter is passed as 'Spanish'", func(t *testing.T) {
		got := Hello("Raphael", "Spanish")
		want := "Hola Raphael"

		assertMessage(t, got, want)
	})
}
