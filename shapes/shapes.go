// Package shapes is package that provides functions for stucts, methods, interfaces  test cases
package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for rectangle: rectangle.Area()
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

// Area method for circle: circle.Area()
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Shape interface {
	Area() float64
}
