package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

// Answer parses and evaluates simple math word problems
// and returns the answer as an integer.
func Answer(question string) (int, bool) {
	t := parseAndSplit(question)

	var n int
	// reduce until length will be 1
	for len(t) > 1 {
		if len(t) < 3 {
			return 0, false
		}

		n, err := eval(t[0], t[1], t[2])
		if err != nil {
			return 0, false
		}

		t = append([]string{strconv.Itoa(n)}, t[3:]...)
	}

	n, _ = strconv.Atoi(t[0])
	return n, true
}

// eval takes the first operand, operator and the second operand,
// and evaluates the result
func eval(first, op, second string) (int, error) {
	a, err := strconv.Atoi(first)
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(second)
	if err != nil {
		return 0, err
	}

	switch op {
	case "minus":
		return a - b, nil
	case "plus":
		return a + b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		return a / b, nil
	default:
		return 0, fmt.Errorf("something wrong")
	}
}

// parse the given string, delete suffix and prefix, split in the slice
func parseAndSplit(s string) []string {
	s = strings.TrimPrefix(s, "What is ")
	s = strings.TrimSuffix(s, "?")
	s = strings.ReplaceAll(s, "multiplied by", "multiply")
	s = strings.ReplaceAll(s, "divided by", "divide")
	return strings.Split(s, " ")
}
