package rotationalcipher

import "strings"

// RotationalCipher implements the rotational cipher. It takes input string
// and return encoded output.
func RotationalCipher(input string, rot int) string {
	var output strings.Builder
	var base int

	for _, char := range input {
		if char < 'A' || char > 'z' || char < 'a' && char > 'Z' {
			output.WriteRune(char)
			continue
		}

		base = 97
		if char < 'a' {
			base = 65
		}

		newChar := (int(char)+rot-base)%26 + base
		output.WriteRune(rune(newChar))
	}

	return output.String()
}
