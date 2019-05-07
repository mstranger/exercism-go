package piglatin

import (
	"fmt"
	"regexp"
)

const addition = "ay"

// Sentence translates a sentence from English to Pig Latin.
func Sentence(input string) (output string) {
	for _, w := range regexp.MustCompile(`[[:space:]]`).Split(input, -1) {
		w, _ = Word(w)
		output += w + " "
	}
	// last char is a space
	output = output[:len(output)-1]
	return
}

// Word translates one word from English to Pig Latin.
func Word(input string) (result string, err error) {
	// return error if a given input is a sentence
	if regexp.MustCompile(`[[:space:]]`).MatchString(input) == true {
		err = fmt.Errorf("invalid input: must be only one word")
		return
	}

	// vowel sound or 'xr', 'yt' at the beginning
	rule1 := regexp.MustCompile(`(^[aeoui])|(^xr.+)|(^yt.+)`)
	// 	a word begins with a consonant sound
	rule2 := regexp.MustCompile(`^[^aeoui]+`)
	// a word starts with a consonant sound followed by 'qu'
	rule3 := regexp.MustCompile(`^(.+)?qu`)
	// a word contains a 'y' after a consonant cluster
	rule4 := regexp.MustCompile(`^[^aeoui]+y`)

	if rule1.FindString(input) != "" {
		result = input + addition
		return
	}

	if chunk := rule3.FindString(input); chunk != "" {
		result = input[len(chunk):] + chunk + addition
		return
	}

	if chunk := rule4.FindString(input); chunk != "" {
		result = input[len(chunk)-1:] + chunk[:len(chunk)-1] + addition
		return
	}

	if chunk := rule2.FindString(input); chunk != "" {
		result = input[len(chunk):] + chunk + addition
		return
	}

	return
}
