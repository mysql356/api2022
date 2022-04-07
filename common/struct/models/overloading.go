package models

import "math"

func (r Reactangle) CalculateArea() int {
	return r.Lenght * r.Width
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}
