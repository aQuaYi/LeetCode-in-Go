package problem0902

import (
	"math"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	ns := strconv.Itoa(N)
	size := len(D)
	res := empty(size, len(ns))
	return res + equal(D, ns)
}

//
func equal(D []string, N string) int {
	if len(N) == 0 {
		return 1
	}
	size := len(D)
	res := 0
	head := N[0:1]
	for _, d := range D {
		if d < head {
			res += less(size, len(N))
		} else if d == head {
			res += equal(D, N[1:])
		}
	}
	return res
}

// size 个数，组合成低于 length 位的数字的组合数
func empty(size, length int) int {
	s := float64(size)
	l := float64(length)
	res := (math.Pow(s, l) - s) / (s - 1)
	return int(res)
}

// size 个数，组成 length-1 位数字的组合数
func less(size, length int) int {
	s := float64(size)
	l := float64(length)
	res := math.Pow(s, l-1)
	return int(res)
}
