package cipher

import (
	"regexp"
	"unicode"
)

// Cipher interface implementations.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// cipher holds the distances
type cipher []int

// Endoce implements encrypt a message.
func (c cipher) Encode(s string) string {
	return c.shiftLetters(s, func(a, b int) int { return a + b })
}

// Decode implements decrypt a message.
func (c cipher) Decode(s string) string {
	return c.shiftLetters(s, func(a, b int) int { return a - b })
}

// shiftLetters encodes a letter based on some function.
func (c cipher) shiftLetters(s string, f func(a, b int) int) string {
	shifted := ""
	for _, v := range s {
		if !unicode.IsLetter(v) {
			continue
		}

		dist := c[len(shifted)%len(c)]
		ch := f(int(unicode.ToLower(v)), dist)
		if ch > 'z' {
			ch -= 'z' - 'a' + 1
		}
		if ch < 'a' {
			ch += 'z' - 'a' + 1
		}

		shifted += string(ch)
	}

	return shifted
}

// NewCaesar creates a cipher with distance = 3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift creates a new shift cipher with given distance.
func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	c := cipher([]int{distance})
	return c
}

// NewVigenere creates a new Vigenere sipher with given key.
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	if regexp.MustCompile(`a`).ReplaceAllString(key, "") == "" {
		return nil
	}

	ints := make([]int, len(key))
	for i, v := range key {
		if !unicode.IsLower(v) || !unicode.IsLetter(v) {
			return nil
		}

		ints[i] = int(v - 'a')
	}

	c := cipher(ints)
	return c
}
