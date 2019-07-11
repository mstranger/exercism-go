package diamond

import (
	"fmt"
	"strings"
)

// Gen takes a letter and outputs a diamond shape.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", fmt.Errorf("invalid input")
	}

	n := int(b - 'A' + 1)
	s := make([]string, 0)

	for i := 0; i < n; i++ {
		var line strings.Builder
		// row length = 2 * n - 1
		for j := 0; j < 2*n-1; j++ {
			if j == n-1-i || j == n-1+i {
				line.WriteByte(byte(int(b) - n + i + 1))
			} else {
				line.WriteByte(' ')
			}
		}
		s = append(s, line.String())
	}

	s = append(s, reversed(s[:len(s)-1])...)

	return strings.Join(s, "\n") + "\n", nil
}

// return reveversed copy of the given slice
func reversed(arr []string) []string {
	res := make([]string, len(arr))
	for i, v := range arr {
		res[len(res)-1-i] = v
	}
	return res
}
