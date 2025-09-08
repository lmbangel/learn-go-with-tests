package main

import "testing"

func TestSlices(t *testing.T) {

	t.Run("Test summing the elements ins a list", func(t *testing.T) {
		sumThis := [5]int{5, 6, 2, 1, 10}
		got := SumSlices(sumThis)
		want := 24

		if got != want {
			t.Errorf("Expecting %d, but received %d", got, want)
		}
	})
}
