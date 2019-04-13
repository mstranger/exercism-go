package etl

import "strings"

// Transform transforms the legacy data format to the new format.
func Transform(input map[int][]string) map[string]int {
	result := map[string]int{}

	for i, line := range input {
		for _, char := range line {
			result[strings.ToLower(char)] = i
		}
	}

	return result
}
