package diffsquares

func SquareOfSum(n int) int {
    sum:= 0
	if n < 1{
        return sum
    }
    sum = n * (n + 1) / 2
    return sum * sum
}

func SumOfSquares(n int) int {
	if n < 1{
        return 0
    }
    return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
