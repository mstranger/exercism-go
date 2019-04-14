package anagram

import (
	"sort"
	"strings"
)

type sortRunes []rune

// satisfy sort.Interface{} for sortRunes type
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Len() int           { return len(s) }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// sort chars in increasing order
func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// Detect selects all possible anagrams from the given list.
func Detect(s string, candidates []string) []string {
	detected := make([]string, 0)
	sortedInput := sortString(strings.ToLower(s))

	for _, word := range candidates {
		if strings.ToLower(word) == strings.ToLower(s) {
			continue
		}
		if sortedInput == sortString(strings.ToLower(word)) {
			detected = append(detected, word)
		}
	}

	return detected
}
