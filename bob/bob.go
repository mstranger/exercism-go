package bob

import "strings"

// Hey provides teenager answers.
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYelling(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isQuestion(s string) bool {
	return s[len(s)-1] == '?'
}

func isYelling(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}
