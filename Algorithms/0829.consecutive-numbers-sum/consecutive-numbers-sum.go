package problem0829

import (
	"math"
)

func consecutiveNumbersSum(N int) int {
	res := 1

	root := int(math.Sqrt(float64(N)))
	for i := 2; i <= root; i++ {
		res += analyse(N, i)
	}

	return res
}

func analyse(n, d int) int {
	q, r := n/d, n%d

	if r == 0 && isOdd(d) {
		if isOdd(q) && d != q && d-q/2 > 0 {
			return 2
		}
		return 1
	}

	if r*2 == d {
		return 1
	}

	return 0
}

func isOdd(i int) bool {
	return i%2 == 1
}
