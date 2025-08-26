package main

import "testing"

func TestAdder(t *testing.T) {
	t.Run("When two integers are passed", func(t *testing.T) {
		got := Adder(2, 2)
		want := 4

		if got != want {
			t.Errorf("Got %q, expecting %q ", got, want)
		}
	})
}
