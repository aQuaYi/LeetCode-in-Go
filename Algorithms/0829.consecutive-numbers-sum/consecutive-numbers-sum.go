package problem0829

import (
	"math"
)

func consecutiveNumbersSum(N int) int {
	res := 1

	root := int(math.Sqrt(float64(2 * N)))
	for k := 2; k <= root; k++ {
		kx := N - k*(k-1)/2
		if kx%k == 0 {
			res++
		}
	}

	return res
}

//
