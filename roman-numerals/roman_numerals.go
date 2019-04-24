package romannumerals

import (
	"fmt"
	"strings"
)

var romanLetters = []string{
	"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
var decimalValues = []int{
	1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

// ToRomanNumeral converts from normal (decimal) numbers to Roman Numerals.
func ToRomanNumeral(arabic int) (string, error) {
	var res strings.Builder

	if arabic < 1 || arabic > 3000 {
		return "", fmt.Errorf("the number should be in the range 1..3000")
	}

	for i := len(decimalValues) - 1; i >= 0; i-- {
		for decimalValues[i] <= arabic {
			res.WriteString(romanLetters[i])
			arabic -= decimalValues[i]
		}
	}

	return res.String(), nil
}
