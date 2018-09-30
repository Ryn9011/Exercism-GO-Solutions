package diffsquares

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

func Difference(num int) int {
	squareOfSum := SquareOfSums(num)
	sumOfSquare := SumOfSquares(num)
	return squareOfSum - sumOfSquare
}
