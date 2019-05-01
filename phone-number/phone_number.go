package phonenumber

import (
	"fmt"
	"strings"
)

// Number returns a string with digits only.
func Number(s string) (string, error) {
	s, err := normalize(s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// AreaCode returns an area code of the given number.
func AreaCode(s string) (string, error) {
	s, err := normalize(s)
	if err != nil {
		return "", err
	}

	return s[:3], nil
}

// Format produces a pretty string from the given number.
func Format(s string) (string, error) {
	s, err := normalize(s)
	if err != nil {
		return "", err
	}

	// adding strings will be more faster
	return fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:]), nil
}

// normalize deletes all non-digits chars and validates given number
func normalize(s string) (string, error) {
	// more concise, but more slowly
	// re := regexp.MustCompile(`[[:^digit:]]`)
	// s = re.ReplaceAllString(s, "")

	filtered := strings.Builder{}
	for _, v := range s {
		if v < '0' || v > '9' {
			continue
		}
		filtered.WriteRune(v)
	}
	s = filtered.String()

	if len(s) < 10 || len(s) > 11 {
		return "", fmt.Errorf("invalid number: length")
	}

	if len(s) == 11 {
		if s[0] != '1' {
			return "", fmt.Errorf("invalid number: must start with 1")
		}
		s = s[1:]
	}

	if s[0] < '2' || s[3] < '2' {
		return "", fmt.Errorf("invalid number: exchange code or area code")
	}

	return s, nil
}
