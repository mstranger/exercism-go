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

	for i, v := range s {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		// double every second digit, starting from the right
		if (len(s)-i)&1 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	return sum%10 == 0
}
