package main

import "testing"

func TestYo(t *testing.T) {
	actual := Yo("Bro")
	expected := "Yo, yoyoyo, Bro!"

	if actual != expected {
		t.Errorf("got %q when %q was expected", actual, expected)
	}
}
