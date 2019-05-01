package sieve

import "math"

// Sieve implements the Sieve of Eratosthenes to find all primes
// from 2 up to a given number.
func Sieve(limit int) []int {
	// 0 - prime
	// 1 - not prime
	sieve := make([]int, limit+1)
	primes := make([]int, 0)

	for i := 0; i <= int(math.Sqrt(float64(limit))); i++ {
		if i < 2 {
			sieve[i] = 1
			continue
		}

		if sieve[i] == 0 {
			for j := i * i; j <= limit; j += i {
				sieve[j] = 1
			}
		}
	}

	for i, v := range sieve {
		if v == 0 {
			primes = append(primes, i)
		}
	}

	return primes
}
