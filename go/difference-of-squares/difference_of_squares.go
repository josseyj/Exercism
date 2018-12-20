//Package diffsquares provides methods of difference of squares
package diffsquares

//SumOfSquares calculates the sum of squares of n numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

//SquareOfSum calculates the squares of sum of n numbers
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

//Difference calculates the difference between the square of sum of numbers
// and the sum of squares of the numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
