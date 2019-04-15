package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if given number is valid or not per the Luhn formula.
func Valid(s string) bool {
	s, err := transformString(s)

	if err != nil || len(s) == 1 {
		return false
	}

	var sum int
	for _, v := range s {
		n, _ := strconv.Atoi(string(v))
		sum += n
	}

	return sum%10 == 0
}

// double every second digit, starting from the right.
// delete all spaces. return error if any non-digit character found.
func transformString(s string) (string, error) {
	s = strings.ReplaceAll(s, " ", "")

	_, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}

	for i := len(s) - 2; i >= 0; i -= 2 {
		n, _ := strconv.Atoi(string(s[i]))
		if n > 9 {
			n = 2*n - 9
		} else {
			n = 2 * n
		}
		s = s[:i] + strconv.Itoa(n) + s[i+1:] // replace s[i]
	}

	return s, nil
}
