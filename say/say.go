package say

import "strings"

const (
	// MinNumber is the minimum number that can be said.
	MinNumber = 0
	// MaxNumber is the maximum number that can be said.
	MaxNumber = 999999999999
)

// Dict contains string equivalents for numbers.
var Dict = map[int64]string{
	0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten", 11: "eleven",
	12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
	17: "seventeen", 18: "eighteen", 19: "ninteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
	80: "eighty", 90: "ninety",
}

var stagesInt = []int64{1000000000, 1000000, 1000, 100, 1}
var stagesString = []string{"billion", "million", "thousand", "hundred", ""}

// Say spells out the given number in English.
func Say(n int64) (string, bool) {
	if n < MinNumber || n > MaxNumber {
		return "", false
	}

	if v, ok := Dict[n]; ok {
		return v, ok
	}

	var s string
	stages := []int64{}
	for _, v := range stagesInt {
		d := n / v
		n = n % v
		stages = append(stages, d)
	}

	// { billion: 1, hundred: 2 ... }
	zip := zipStage(stagesString, stages)

	for i, v := range stagesString {
		t, _ := saySimple(zip[v])
		if t == "zero" {
			continue
		}

		s += strings.TrimSpace(t+" "+stagesString[i]) + " "
	}

	return strings.TrimSpace(s), true
}

// spell out the number in range 0..999
func saySimple(n int64) (string, bool) {
	if n > 999 {
		return "", false
	}

	if v, ok := Dict[n]; ok {
		return v, ok
	}

	s := ""
	if r := n / 100; r != 0 {
		s += Dict[r] + " hundred"
	}

	if last := sayLastTwoDigits(n); last != "" {
		s += " " + last
	}

	return s, true
}

// return the "count" of each stage
func zipStage(s1 []string, s2 []int64) map[string]int64 {
	zipped := map[string]int64{}
	for i, v := range s1 {
		zipped[v] = s2[i]
	}
	return zipped
}

// spell out two last digit for the number
func sayLastTwoDigits(n int64) string {
	n = n % 100

	if n == 0 {
		return ""
	}

	if v, ok := Dict[n]; ok {
		return v
	}

	return Dict[n/10*10] + "-" + Dict[n%10]
}
