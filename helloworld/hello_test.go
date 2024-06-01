package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello with a name but without language", func(t *testing.T) {
		got := Hello("Bro", "")
		want := "Hello, Bro"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello without providing a name and language", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jo", "French")
		want := "Bonjour, Jo"
		assertCorrectMessage(t, got, want)
	})
}

// Helper message for asserting correct messages
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
