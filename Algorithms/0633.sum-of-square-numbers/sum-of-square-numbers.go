package problem0633

import (
	"math"
)

func judgeSquareSum(c int) bool {
	a := intSqrt(c)
	for a >= 0 {
		if isSquare(c - a*a) {
			return true
		}
		a--
	}
	return false
}

func intSqrt(c int) int {
	return int(math.Sqrt(float64(c)))
}

func isSquare(b2 int) bool {
	b := intSqrt(b2)
	return b*b == b2
}
