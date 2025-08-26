package main

import "testing"

func TestRepeater(t *testing.T) {
	t.Run("When a value is passed", func(t *testing.T) {
		got := Repeater("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q, expecting %q", got, want)
		}
	})

	t.Run("When 0 is passed as the iterator ", func(t *testing.T) {
		got := Repeater("f", 0)
		want := ""

		if got != want {
			t.Errorf("got %q, expecting %q", got, want)
		}
	})
}

func BenchmarkRepeater(b *testing.B) {
	for b.Loop() {
		Repeater("a", 5)
	}

}
func BenchmarkRepeater2(b *testing.B) {
	for b.Loop() {
		Repeater2("a", 5)
	}
}
