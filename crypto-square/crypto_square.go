package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode implements the classic method for composing secret messages
// called a square code.
func Encode(input string) string {
	var cipher strings.Builder
	// normalize
	input = strings.ToLower(
		regexp.MustCompile("[[:^word:]]").ReplaceAllString(input, ""))

	// transpose r and c in the output square
	c, r := plainSquareSides(len(input)) // row, column

	for i := 0; i < r; i++ {
		if i != 0 {
			cipher.WriteByte(' ')
		}
		for j := 0; j < c; j++ {
			cur := i + j*r
			if cur < len(input) {
				cipher.WriteByte(input[cur])
			} else {
				cipher.WriteByte(' ')
			}
		}
	}

	return cipher.String()
}

// plainSquareSides returns the sides of the input crypto square.
func plainSquareSides(square int) (int, int) {
	side := int(math.Sqrt(float64(square)))

	// 25 -> 5 x 5
	if side*side == square {
		return side, side
	}

	// 33 -> 6 x 6
	if side*(side+1) < square {
		return side + 1, side + 1
	}

	// other
	return side, side + 1
}
