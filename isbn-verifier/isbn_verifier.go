package isbn

import (
	"fmt"
	"strconv"
	"strings"
)

// IsValidISBN checks if a given string is a valid ISBN number.
func IsValidISBN(isbn string) bool {
	isbn, err := normalizeISBN(isbn)
	if err != nil {
		return false
	}

	var sum int
	for i, v := range isbn {
		var t int
		if v == 'X' {
			t = 10
		} else {
			t = int(v - '0')
		}
		sum += t * (10 - i)
	}

	return sum%11 == 0
}

// normalizeISBN deletes dashes, checks length, checks for the presence
// of a non-digit character in the middle
func normalizeISBN(s string) (string, error) {
	s = strings.ReplaceAll(s, "-", "")

	if len(s) < 10 {
		return "", fmt.Errorf("too short isbn")
	}
	if len(s) > 10 {
		return "", fmt.Errorf("too long isbn")
	}

	t := s
	if strings.HasSuffix(t, "X") {
		t = t[:len(s)-1]
	}

	_, err := strconv.Atoi(t)
	if err != nil {
		return "", err
	}

	return s, nil
}
