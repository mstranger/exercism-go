package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode implements the classic method for composing secret messages
// called a square code.
func Encode(input string) string {
	if input == "" {
		return ""
	}

	chunks := make([]string, 0)
	encoded := make([]string, 0)
	var output string

	// normalize
	re := regexp.MustCompile(`[^a-z0-9]`)
	output = re.ReplaceAllLiteralString(strings.ToLower(input), "")

	// "abcdefgh" -> ["abc" "def" "gh"]
	r, c := squareSides(len(output)) // row, column
	for i := 0; i < r*c; i += c {
		if i+c < len(output) {
			chunks = append(chunks, output[i:i+c])
		} else {
			chunks = append(chunks, output[r*c-c:])
		}
	}

	// add spaces to the last chunk, "abc" -> ["ab", "c "]
	last := chunks[len(chunks)-1]
	if len(last) < c {
		chunks[len(chunks)-1] = chunks[len(chunks)-1] + strings.Repeat(" ", c-len(last))
	}

	// encode
	for j := 0; j < c; j++ {
		s := ""
		for i := 0; i < r; i++ {
			s += string(chunks[i][j])
		}
		encoded = append(encoded, s)
	}

	return strings.Join(encoded, " ")
}

// squareSides returns the sides of the crypto square.
func squareSides(square int) (int, int) {
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
