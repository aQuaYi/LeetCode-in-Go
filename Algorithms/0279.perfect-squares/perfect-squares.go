package Problem0279

import (
	"math"
)

func numSquares(n int) int {
	if n < 4 {
		return n
	}

	temp := int(math.Sqrt(float64(n)))

	return n/temp + numSquares(n%temp)
}
