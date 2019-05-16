package brackets

var matched = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

// Bracket verifies that any and all pairs are matched and nested correctly.
func Bracket(input string) bool {
	stack := make([]rune, 0)

	for _, s := range input {
		// skip any other symbol
		if s != '{' && s != '}' && s != '[' && s != ']' &&
			s != '(' && s != ')' {
			continue
		}

		if v, ok := matched[s]; ok {
			stack = append(stack, v)
		} else {
			// example: "]"
			if len(stack) == 0 {
				return false
			}
			e := stack[len(stack)-1]
			// if current s doesn't match with last e in stack
			if e != s {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// if stack is not empty
	if len(stack) > 0 {
		return false
	}

	return true
}
