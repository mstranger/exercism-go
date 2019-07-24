package transpose

import "strings"

// TransposeV1 transposes the given text.
func TransposeV1(input []string) []string {
	output := make([]string, 0)
	max := maxLength(input)

	line := make([]string, len(input))
	for i := 0; i < max; i++ {
		for j, word := range input {
			if len(word) > i {
				line[j] = string(word[i])
			}
		}
		// line will be something like {"", "", "a", "", "b", "c", "", ""}
		padSpaceLeft(&line)
		output = append(output, strings.Join(line, ""))
		line = make([]string, len(input))
	}

	return output
}

// padSpaceLeft changes all "" to spaces to the left in the given string
func padSpaceLeft(s *[]string) {
	var i int
	for i = len(*s) - 1; i >= 0; i-- {
		if (*s)[i] != "" {
			break
		}
	}

	for j := 0; j < i; j++ {
		if (*s)[j] == "" {
			(*s)[j] = " "
		}
	}
}

// find the longest word in the given input
func maxLength(input []string) (max int) {
	for _, v := range input {
		if max < len(v) {
			max = len(v)
		}
	}
	return
}
