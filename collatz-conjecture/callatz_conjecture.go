package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	var step int

	if n <= 0 {
		return 0, fmt.Errorf("should be positive number")
	}

	for n > 1 {
		if n&1 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}

		step++
	}

	return step, nil
}
