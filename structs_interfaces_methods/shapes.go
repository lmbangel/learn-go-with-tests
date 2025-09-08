package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (s *Rectangle) Area() float64 {
	return s.Width * s.Height
}

type Circle struct {
	Radius float64
}

func (s *Circle) Area() float64 {
	return (s.Radius * s.Radius) * math.Pi
}

func Perimeter(s Rectangle) float64 {
	perimeter := 2 * (s.Width + s.Height)
	return perimeter
}
