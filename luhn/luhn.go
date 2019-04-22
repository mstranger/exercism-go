package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if given number is valid or not per the Luhn formula.
func Valid(s string) bool {
	var sum int

	s = strings.ReplaceAll(s, " ", "")

	if len(s) == 1 {
		return false
	}

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}

		// every second, starting from the right
		if j&1 == 1 {
			if 2*n > 9 {
				n = 2*n - 9
			} else {
				n = 2 * n
			}
		}

		sum += n
	}

	return sum%10 == 0
}
