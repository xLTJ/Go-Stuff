package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers (0.2657 ns/op)
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first N natural numbers (0.2411 ns/op)
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
