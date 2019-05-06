package prime

// Factors computes all prime factors of the given number.
func Factors(input int64) []int64 {
	primes := make([]int64, 0)
	i := int64(2)

	for ; input > 1 && input >= i*i; i++ {
		for input%i == 0 {
			primes = append(primes, i)
			input /= i
		}
	}

	if input > 1 {
		primes = append(primes, input)
	}

	return primes
}
