package railfence

import (
	"strings"
	"unicode"
)

// TODO: use Interfaces?

// Encode implements encoding for the rail fence cipher.
func Encode(text string, rails int) string {
	// delete all punctuation and spaces
	text = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return -1
		}
		return r
	}, text)

	chunks := make([]strings.Builder, rails)

	pos, direction := 0, 1
	for _, c := range text {
		chunks[pos].WriteRune(c)
		pos += direction
		if pos == rails-1 || pos == 0 {
			direction = -direction
		}
	}

	var encoded string
	for _, chunk := range chunks {
		encoded += chunk.String()
	}

	return encoded
}

// Decode implements decoding for the rail fence cipher.
func Decode(input string, rails int) string {
	// find the lengths of all chunks
	chunkLens := make([]int, rails)
	pos, direction := 0, 1
	for i := 0; i < len(input); i++ {
		chunkLens[pos]++
		pos += direction
		if pos == rails-1 || pos == 0 {
			direction = -direction
		}
	}

	// find each chunk string
	chunks, start := make([]string, len(chunkLens)), 0
	for i, v := range chunkLens {
		chunks[i] = input[start : start+v]
		start += v
	}

	// decode input string
	result := ""
	pos, direction = 0, 1
	for i := 0; i < len(input); i++ {
		// get first letter and shift
		result += string(chunks[pos][0])
		chunks[pos] = chunks[pos][1:]
		pos += direction

		if pos == rails-1 || pos == 0 {
			direction = -direction
		}
	}

	return result
}
