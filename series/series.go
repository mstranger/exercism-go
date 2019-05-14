package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	if n > len(s) {
		return []string{}
	}

	result := make([]string, 0)

	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}

	return result
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if res, ok := First(n, s); ok {
		return res
	}

	panic(nil)
}

// First returns the first substring of s and boolean value.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}

	first = s[:n]
	ok = true
	return
}
