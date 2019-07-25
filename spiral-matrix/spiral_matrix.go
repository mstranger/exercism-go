package spiralmatrix

// SpiralMatrix builds a square matrix (n x n) of numbers in spiral order.
func SpiralMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	fillSpiral(res, 1)

	return res
}

// fillSpiral fills the given matrix with numbers
func fillSpiral(m [][]int, s int) int {
	current := s
	n := len(m)

	for t := n; t > n/2; t-- {
		for j := n - t; j < t; j++ {
			m[n-t][j] = current
			current++
		}

		for i := n - t + 1; i < t; i++ {
			m[i][t-1] = current
			current++
		}

		for j := t - 2; j > n-t; j-- {
			m[t-1][j] = current
			current++
		}

		for i := t - 1; i > n-t; i-- {
			m[i][n-t] = current
			current++
		}
	}

	return current
}
