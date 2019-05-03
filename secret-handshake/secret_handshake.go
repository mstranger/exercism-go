package secret

const (
	wink uint = 1 << iota
	doubleBlink
	closeYourEyes
	jump
	reversed
)

// Handshake converts a dicimal number to the appropriate sequence of events.
func Handshake(n uint) []string {
	result := make([]string, 0)

	if n&wink != 0 {
		result = append(result, "wink")
	}
	if n&doubleBlink != 0 {
		result = append(result, "double blink")
	}
	if n&closeYourEyes != 0 {
		result = append(result, "close your eyes")
	}
	if n&jump != 0 {
		result = append(result, "jump")
	}
	if n&reversed != 0 {
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
