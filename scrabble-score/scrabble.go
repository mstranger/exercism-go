package scrabble

import "strings"

var values = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score computes the scrabble score for the given word.
func Score(input string) int {
	points := 0
	for _, char := range input {
		for k, v := range values {
			if strings.Index(k, strings.ToUpper(string(char))) != -1 {
				points += v
				break
			}
		}
	}

	return points
}
