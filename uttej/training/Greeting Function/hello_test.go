package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {

		t.Helper() /* this tells the compiler that the function is a
		helper function and when the test fails, it return the line number where i
		t found the flaw instead of sending the line number inside this function which
		wouldn't be of much use. */
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Uttej", "")
		want := "Hello Uttej"

		// if got != want { this is getting refactored

		// 	t.Errorf("got %q, want %q", got, want)
		// }
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello world when empty string is passed", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {

		got := Hello("Iniesta", "Spanish")
		want := "Hola Iniesta"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French", func(t *testing.T) {

		got := Hello("Kylian", "French")
		want := "Bonjour Kylian"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in telugu", func(t *testing.T) {

		got := Hello("Uttej", "Telugu")
		want := "Namaskaram Uttej"

		assertCorrectMessage(t, got, want)
	})

}
