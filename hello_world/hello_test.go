package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("When a name is passed.", func(t *testing.T) {
		name := "Lihle"
		g := SayHello(name)
		w := "Hello, Lihle"
		
		assertCorrectMessage(t, g, w)
	})

	t.Run("When no name is not passed.", func(t *testing.T) {
		name := ""
		g := SayHello(name)
		w := "Hello, World"

		assertCorrectMessage(t, g, w)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q.", got, want)
	}
}
