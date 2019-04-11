package triangle

import "math"

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides determine if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	// if any side is NaN
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	// if any side is Inf
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	// if any side <= 0
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	// if the shape is not a triangle
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}

	// default `scalene`
	return Sca
}
