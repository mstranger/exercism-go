package wordcount

import (
	"regexp"
	"strings"
)

// Frequency type contains the frequency of each word.
type Frequency map[string]int

// WordCount counts the occurrences of each word in phrase.
func WordCount(phrase string) Frequency {
	result := Frequency{}

	// split ` '` or `' ` but not `a'b`
	r := regexp.MustCompile(`'\s|\s'|[^\w']`)

	for _, word := range r.Split(phrase, -1) {
		if word == "" {
			continue
		}
		result[strings.ToLower(word)]++
	}

	return result
}
