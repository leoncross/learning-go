package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Leon", "")
		want := "Hello, Leon"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Leon", "Spanish")
		want := "Hola, Leon"
		assertCorrectMessage(t, got, want)
	})

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Leon", "French")
    want := "Bonjour, Leon"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Swahili", func(t *testing.T) {
    got := Hello("Leon", "Swahili")
    want := "Jambo, Leon"
    assertCorrectMessage(t, got, want)
  })

}
