package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.breadth + r.length)
}

func (r Rectangle) Area() float64 {
	return (r.breadth * r.length)
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * c.radius
}
