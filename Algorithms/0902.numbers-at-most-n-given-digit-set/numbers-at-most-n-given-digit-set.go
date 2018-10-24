package problem0902

import (
	"math"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	ns := strconv.Itoa(N)
	return less(D, ns)
}

// 更高位小于 N 的最高位
func less(D []string, N string) int {
	a := float64(len(D))
	s := float64(len(N))
	res := (math.Pow(a, s) - a) / (a - 1)
	return int(res)
}

// 更高位等于 N 的最高位
func equal(D []string, N string) int {
	i := 0
	for i < len(D) && D[i] < N[1] {

	}
	return 0
}
