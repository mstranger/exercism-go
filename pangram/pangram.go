package pangram

import (
	"strings"
	"unicode"
)

// N is the number of letters in the (English) alphabet.
const N int = 26

// IsPangram determines if the sentence is a pangram.
func IsPangram(s string) bool {
	hash := map[rune]int{}
	for _, v := range strings.ToLower(s) {
		if !unicode.IsLetter(v) {
			continue
		}
		hash[v]++
	}

	if len(hash) < N {
		return false
	}

	return true
}
