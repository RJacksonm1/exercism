package diffsquares

// SquareOfSums 🤯
func SquareOfSums(n int) int {
	sumOfNaturals := (n * (n + 1)) / 2
	return sumOfNaturals * sumOfNaturals
}

// SumOfSquares 🤯
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference 🤯
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
