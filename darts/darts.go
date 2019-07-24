package darts

import "math"

// Score returnes earned points in a single toss of a Darts game.
func Score(x, y float64) int {
	points := math.Sqrt(x*x + y*y)
	switch {
	case points <= 1.0:
		return 10
	case points <= 5.0:
		return 5
	case points <= 10.0:
		return 1
	default:
		return 0
	}
}
