package ocr

import "strings"

// Recognize determines which number is represented in the given
// a 3 x 4 grid of pipes.
// If the input is incorrect or not recognized it will return "?" sign.
func Recognize(s string) []string {
	r := make([]string, 0)

	line := "\n"
	for i, v := range strings.Split(s[1:], "\n") {
		line += v + "\n"
		if (i+1)%4 == 0 {
			r = append(r, recognizeNumber(line[:len(line)-1]))
			line = "\n"
		}
	}

	return r
}

// recognizeDigit recognizes one digit from 0 to 9
func recognizeDigit(s string) string {
	r := make([]string, 0)

	for _, v := range strings.Split(s[1:], "\n") {
		// grid 3x4
		if len(v) != 3 {
			return ""
		}

		r = append(r, v)
	}

	// grid 3x4
	if len(r) != 4 {
		return ""
	}

	if r[0] == "   " {
		if r[1] == "  |" {
			return "1"
		}
		if r[1] == "|_|" {
			return "4"
		}
	}

	if r[1] == "  |" {
		return "7"
	}

	if r[2] == "|_ " {
		return "2"
	}

	if r[1] == " _|" {
		return "3"
	}

	if r[1] == "|_ " {
		if r[2] == " _|" {
			return "5"
		}
		return "6"
	}

	if r[0] == " _ " && r[1] == "| |" && r[2] == "|_|" {
		return "0"
	}

	if r[1] == "|_|" && r[2] == "|_|" {
		return "8"
	}

	if r[1] == "|_|" && r[2] == " _|" {
		return "9"
	}

	// undefined digit
	return "?"
}

// recognizeNumber recognizes the given number
func recognizeNumber(s string) string {
	r := splitDigits(s)
	digits := make([]string, len(r))
	for i, v := range r {
		digits[i] = recognizeDigit(v)
	}

	return strings.Join(digits, "")
}

// splitDigits splits all digits in the given number
func splitDigits(s string) []string {
	r := make([]string, 0)

	sl := strings.Split(s, "\n")
	cnt := len(sl[1]) / 3

	for i := 0; i < cnt; i++ {
		l := "\n"
		for _, line := range sl[1:] {
			l += line[i*3:i*3+3] + "\n"
		}
		r = append(r, l[:len(l)-1])
	}

	return r
}

const TestVersion = 1
