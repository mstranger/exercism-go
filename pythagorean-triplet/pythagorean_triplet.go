package pythagorean

import "math"

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max
func Range(min, max int) []Triplet {
	tripletList := make([]Triplet, 0)

	for a := min; a <= max-1; a++ {
		for b := a + 1; b <= max; b++ {
			c := math.Sqrt(float64(a*a) + float64(b*b))
			if c <= float64(max) && c == math.Floor(c) {
				t := [3]int{a, b, int(c)}
				tripletList = append(tripletList, t)
			}
		}
	}

	return tripletList
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c = p
func Sum(p int) []Triplet {
	result := make([]Triplet, 0)
	tripletsList := Range(1, p)

	for _, t := range tripletsList {
		sum := t[0] + t[1] + t[2]
		if sum == p {
			result = append(result, t)
		}
	}

	return result
}
