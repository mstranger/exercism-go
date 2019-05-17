package brackets

var matched = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

// Stack represents a rune container with a LIFO rule.
type Stack []rune

// Push adds an element to the stack.
func (s *Stack) Push(e rune) {
	*s = append(*s, e)
}

// Pop gets the last item from the stack and returns it.
func (s *Stack) Pop() rune {
	var e rune
	if len(*s) > 0 {
		e = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
	}
	return e
}

// Bracket verifies that any and all pairs are matched and nested correctly.
func Bracket(input string) bool {
	stack := Stack{}

	for _, s := range input {
		// skip any other symbol
		if s != '{' && s != '}' && s != '[' && s != ']' &&
			s != '(' && s != ')' {
			continue
		}

		if v, ok := matched[s]; ok {
			stack.Push(v)
		} else {
			if e := stack.Pop(); e != s {
				return false
			}
		}
	}

	// if stack is not empty
	if len(stack) > 0 {
		return false
	}

	return true
}
