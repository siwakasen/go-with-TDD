// Package shapes is package that provides functions for stucts, methods, interfaces  test cases
package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}
type Circle struct {
	Radius float64
}

// Area method for rectangle: r.Area()
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Area method for circle: c.Area()
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area method for triangle: t.Area()
func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

// Shape interface with Area
type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
