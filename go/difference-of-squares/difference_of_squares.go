package diffsquares

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Faster than math.Pow()
func SquareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
