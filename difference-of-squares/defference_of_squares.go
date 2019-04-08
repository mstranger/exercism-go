package diffsquares

// SquareOfSum finds the square of the first n numbers.
func SquareOfSum(n int) int {
	s := n * (1 + n) / 2
	return s * s
}

// SumOfSquares finds the sum of the squares of the first n numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference finds the difference between the square of the sum of first n
// numbers and the sum of the squares of the first n numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
