package armstrong

import "strconv"

// IsNumber checks than the number is an Armstrong number.
func IsNumber(input int) bool {
	sum := 0
	s := strconv.Itoa(input)
	factor := len(s)

	for _, c := range s {
		n := int(c) - '0'
		sum += pow(n, factor)
	}

	return sum == input
}

// computes x**y
func pow(x, y int) int {
	var res int = 1
	for ; y > 0; y-- {
		res *= x
	}
	return res
}
