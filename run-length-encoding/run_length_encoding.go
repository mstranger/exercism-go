package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode replaces consecutive data elements by just one data
// value and count ("wwwbba" -> "3w2ba")
func RunLengthEncode(input string) string {
	var output string
	var i, j int

	for ; i < len(input); i += j {
		// find count number
		for j = 0; i+j < len(input) && input[i] == input[i+j]; j++ {
		}

		chunck := input[i : i+j]

		if len(chunck) > 1 {
			s := strconv.Itoa(len(chunck)) + chunck[0:1]
			output += s
		} else {
			output += chunck
		}
	}

	return output
}

// RunLengthDecode replaces count and value by consecutive data elements
// ("3w2ba" -> "wwwbba")
func RunLengthDecode(input string) string {
	var output string
	var j int

	for i := 0; i < len(input); {
		if unicode.IsNumber(rune(input[i])) {
			// find numuber chunck
			for j = 0; i+j < len(input)-1 && unicode.IsNumber(rune(input[j+i])); j++ {
			}

			// convert to int and ignore all errors
			n, _ := strconv.Atoi(input[i : i+j])

			// because 3ba and b should be skipped
			i += j + 1
			output += strings.Repeat(string(input[i-1]), n)
		} else {
			output += input[i : i+1]
			i++
		}
	}

	return output
}
