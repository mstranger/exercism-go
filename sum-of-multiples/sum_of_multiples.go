package summultiples

// SumMultiples finds the sum of all the unique multiples
// of particular numbers up to but not including given number.
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	h := make(map[int]bool, 0)

	for _, v := range divisors {
		if v == 0 {
			continue
		}
		for i := 1; v*i < limit; i++ {
			if h[v*i] {
				continue
			}
			sum += v * i
			h[v*i] = true
		}
	}

	return sum
}
