package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	// to upper, delete `'` and `,`, change `-` to space
	replacer := strings.NewReplacer("-", " ", ",", "", "'", "")
	s = replacer.Replace(strings.ToUpper(s))

	// match the first letter of each word
	re := regexp.MustCompile(`\b.`)
	abbr := make([]byte, 0)

	for _, v := range re.FindAll([]byte(s), -1) {
		if string(v) == " " {
			continue
		}
		abbr = append(abbr, v...)
	}

	return string(abbr)
}
