package diffsquares
import "math"

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func SquareOfSum(n int) int {
    sos := 0
	for i := 1; i <= n; i++{
       sos += i
    }
    return int(powInt(sos,2))
}

func SumOfSquares(n int) int {
    sof := 0
	for i := 1; i <= n; i++{
       sof += (powInt(i,2))
    }
    return sof
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

