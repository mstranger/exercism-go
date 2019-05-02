package secret

// decimal equivalents of the 1, 10, 100, 1000 and 10000 in binary
var codes = [5]uint{1, 2, 4, 8, 16}
var values = [5]string{"wink", "double blink", "close your eyes", "jump", ""}

// Handshake converts a dicimal number to the appropriate sequence of events.
func Handshake(n uint) []string {
	result := make([]string, 0)

	for i, v := range codes[:4] {
		// for example:
		// 1111 & 1000 = 1000 in binary; 15 & 8 == 8 in decimal
		if n&v == 0 {
			continue
		}

		result = append(result, values[i])
	}

	// reverse array if addition of 16
	if n&codes[4] != 0 {
		reverse(result)
	}

	return result
}

// reverse the array of strings
func reverse(input []string) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
