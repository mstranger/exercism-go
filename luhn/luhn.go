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

	double := len(s)%2 == 0
	for _, v := range s {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		double = !double
		sum += n
	}

	return sum%10 == 0
}
