package prime

import "math"

// Nth calculates n-th prime number.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	primes := []int{2}
	candidate := 3

	for len(primes) < n {
		if isPrimeWithPrev(candidate, primes) {
			primes = append(primes, candidate)
		}

		// 4, 6, 8 ... are not prime by default
		candidate += 2
	}

	return primes[len(primes)-1], true
}

// if given number isn't prime, it's divided by some previous prime number
func isPrimeWithPrev(n int, primes []int) bool {
	if n < 2 {
		return false
	}

	limit := math.Sqrt(float64(n)) + 1

	// the divider must be in the primes array
	for i := 0; primes[i] < int(limit); i++ {
		if n%primes[i] == 0 {
			return false
		}
	}

	return true
}
