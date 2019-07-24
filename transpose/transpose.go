package transpose

// Transpose transposes the given text.
func Transpose(input []string) []string {
	max := maxWordLength(input)
	output := make([]string, max)

	for i, row := range input {
		for j, col := range row {
			output[j] += string(col)
		}

		remMax := maxWordLength(input[i:])
		for j := len(row); j < remMax; j++ {
			output[j] += " "
		}
	}

	return output
}

// find the longest word in the given input
func maxWordLength(input []string) (max int) {
	for _, v := range input {
		if max < len(v) {
			max = len(v)
		}
	}
	return
}

const testVersion = 2
