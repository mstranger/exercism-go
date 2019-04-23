package grains

import "fmt"

// Square calculates the number of grains on a given square of a chessboard.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("n must be between 1 and 64")
	}

	return 1 << uint(n-1), nil
}

// Total calculates the number of grains on the entire chessboard.
func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		t, _ := Square(i)
		sum += t
	}
	return sum
}
