package diffsquares

// SquareOfSum iterates from 0 to a given number and returns the square of the sum.
func SquareOfSum(n int) int {
  result := n * (n + 1) / 2
  return result * result
}

// SumOfSquares iterates from 0 to a given number and returns the sum of squares.
func SumOfSquares(n int) int {
  return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns a difference between SquareOfSum and SumOfSquares functions.
func Difference(n int) int {
  return SquareOfSum(n) - SumOfSquares(n)
}
