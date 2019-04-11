package reverse

// String provides reversed sentense.
func String(s string) string {
	reversed := make([]byte, len(s))
	j := len(s)

	for i, v := range s {
		rlen := len(string(v))               // rune length
		copy(reversed[j-rlen:], s[i:i+rlen]) // copy rune from src to dest
		j -= rlen
	}

	return string(reversed)
}
