package problem0837

import (
	"math"
)

func new21Game(N int, K int, W int) float64 {
	var res float64
	search(0, 0, N, K, W, &res)
	return res
}

func search(start, exp, n, k, w int, res *float64) {
	if start >= k {
		if start <= n {
			*res += math.Pow(float64(w), float64(exp))
		}
		return
	}

	for i := 1; i <= w; i++ {
		search(start+i, exp-1, n, k, w, res)
	}
}
