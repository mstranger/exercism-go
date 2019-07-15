package ocr

import "strings"

var digitsMap = map[string]string{
	` _ | ||_|   `: "0",
	`     |  |   `: "1",
	` _  _||_    `: "2",
	` _  _| _|   `: "3",
	`   |_|  |   `: "4",
	` _ |_  _|   `: "5",
	` _ |_ |_|   `: "6",
	` _   |  |   `: "7",
	` _ |_||_|   `: "8",
	` _ |_| _|   `: "9",
}

// Recognize determines which number is represented in the given
// a 3 x 4 grid of pipes.
// If the input is incorrect or not recognized it will return "?" sign.
func Recognize(s string) []string {
	r := make([]string, 0)

	sl := strings.Split(s, "\n")

	line := ""
	for i, v := range sl[1:] {
		line += v + "\n"
		if (i+1)%4 == 0 {
			r = append(r, recognizeNumber(line[:len(line)-1]))
			line = ""
		}
	}

	return r
}

// recognizeDigit recognizes one digit from 0 to 9
func recognizeDigit(s string) string {
	if digit, ok := digitsMap[s]; ok {
		return digit
	}

	return "?"
}

// recognizeNumber recognizes the given number
func recognizeNumber(s string) string {
	sl := strings.Split(s, "\n")
	cnt := len(sl[0]) / 3
	n := make([]string, cnt)

	d := ""
	for i := 0; i < cnt; i++ {
		for j := 0; j < 4; j++ {
			d += sl[j][i*3 : i*3+3]
		}
		n = append(n, recognizeDigit(d))
		d = ""
	}

	return strings.Join(n, "")
}

const TestVersion = 2
