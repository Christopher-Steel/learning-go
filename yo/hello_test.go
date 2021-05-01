package main

import "testing"

func TestYo(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("got %q when %q was expected", actual, expected)
		}
	}

	t.Run("saying yo to a bro", func(t *testing.T) {
		actual := Yo("Bro", "")
		expected := "Yo, yoyoyo, Bro!"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying yo to noone in particular", func(t *testing.T) {
		actual := Yo("", "")
		expected := "Yo, yoyoyo, man!"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying yo in Spanish", func(t *testing.T) {
		actual := Yo("", "spanish")
		expected := "Jo, jojojo, hombre!"
		assertCorrectMessage(t, actual, expected)
	})
}
