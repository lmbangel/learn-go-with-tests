package main

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("Is divisible by 3 & 5", func(t *testing.T) {
		got := FizzBuzz(5)
		want := []string{"1", "2", "Fizz", "4", "Buzz"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q, expecting %q", got, want)
		}
	})
}
