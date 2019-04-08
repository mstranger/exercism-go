package isogram

import "unicode"

// IsIsogram checks if word or phrase is an isogram.
func IsIsogram(s string) bool {
	var hash = map[rune]int{}

	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}

		if hash[unicode.ToLower(c)] == 1 {
			return false
		}

		hash[unicode.ToLower(c)]++
	}

	return true
}
