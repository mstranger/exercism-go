package atbash

import "strings"

const plain = "abcdefghijklmnopqrstuvwxyz0123456789"  // plain text
const cipher = "zyxwvutsrqponmlkjihgfedcba0123456789" // substitute

// Atbash encodes the input string using the atbash cipher.
func Atbash(input string) string {
	var encoded strings.Builder

	spaceIdx := 0 // split encoded with spaces
	for _, v := range strings.ToLower(input) {
		j := strings.IndexRune(plain, v)
		if j == -1 {
			continue
		}

		if spaceIdx > 0 && spaceIdx%5 == 0 {
			encoded.WriteByte(' ')
		}
		encoded.WriteByte(cipher[j])
		spaceIdx++
	}

	return encoded.String()
}
