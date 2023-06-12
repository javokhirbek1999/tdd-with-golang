package main

import "testing"

func TestHello(t *testing.T) {

	// Subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Javokhirbek", "")
		want := "Hello, Javokhirbek"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Javokhirbek", "Spanish")
		want := "Hola, Javokhirbek"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {

		got := Hello("Javokhirbek", "French")
		want := "Bonjour, Javokhirbek"

		assertCorrectMessage(t, got, want)

	})

}

// testing.TB calls both testing.T and benchmark testing.B
func assertCorrectMessage(t testing.TB, got, want string) {

	// Helper function that makes this function a helper
	// so the code nows this function is a helper function

	// Instead of complaining about the error in this function,
	// It references the error that occured in the actual failing test function
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
