package accumulate

// Accumulate implements accumulate operation.
func Accumulate(s []string, f func(string) string) []string {
	result := make([]string, 0)
	for _, v := range s {
		result = append(result, f(v))
	}

	return result
}
