package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Test Perimeter", func(t *testing.T) {
		rectange := Rectangle{10.0, 10.0}
		got := Perimeter(rectange)
		want := 40.0

		if got != want {
			t.Errorf("Got %.2f , expecting %.2f ", got, want)
		}
	})
}

func CheckArea(t testing.TB, s Shape, want float64) {
	t.Helper()
	got := s.Area()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	t.Run("Test the Area of a Rectangle", func(t *testing.T) {
		rectangle := &Rectangle{10.0, 10.0}
		want := 100.0
		CheckArea(t, rectangle, want)
	})

	t.Run("Test the Area of a Circle", func(t *testing.T) {
		circle := &Circle{5.0}
		want := (circle.Radius * circle.Radius) * math.Pi
		CheckArea(t, circle, want)
	})

	t.Run("Test the Area of a Triange", func(t *testing.T) {
		triangle := &Triangle{10.0, 5}
		want := (triangle.Base * triangle.Height) / 2
		CheckArea(t, triangle, want)
	})
}
