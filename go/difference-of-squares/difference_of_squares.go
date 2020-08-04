package diffsquares

import "math"

//SumOfSquares computes the sum of consecutive squares upto n
func SumOfSquares(n int) int {
	return ((2*n + 1) * n * (n + 1)) / 6
}

//SquareOfSum calculates the sum of consecutive numbers upto n squared
func SquareOfSum(n int) int {
	return int(math.Pow(float64(((n * (n + 1)) / 2)), 2))
}

//Difference calculates the difference betweeen SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
