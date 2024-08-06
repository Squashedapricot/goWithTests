package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, got, want float64) {
		if got != want {
			t.Helper()
			t.Errorf("expected %g got %g", want, got)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		checkPerimeter(t, got, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 2 * math.Pi * 10.0
		if got != want {
			t.Errorf("expected %g got %g", want, got)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{length: 10.0, breadth: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g want %g", got, tt.hasArea)
		}
	}
}
