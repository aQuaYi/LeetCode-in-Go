package Problem0564

import (
	"strconv"
)

func nearestPalindromic(n string) string {
	size := len(n)
	candidates := getBitChangeCandidates(size)
	candidates = append(candidates, getCanidates(n)...)

	num, _ := strconv.Atoi(n)
	delta := func(x int) int {
		if x > num {
			return x - num
		}
		return num - x
	}

	res := 1<<63 - 1
	for _, cand := range candidates {
		if cand == num {
			continue
		}

		if delta(cand) < delta(res) ||
			(delta(cand) == delta(res) && cand < res) {
			res = cand
		}
	}

	return strconv.Itoa(res)
}

func getCanidates(n string) []int {
	size := len(n)
	prefix := n[:(size+1)/2]
	p, _ := strconv.Atoi(prefix)

	res := make([]int, 3)
	res[0] = p - 1
	res[1] = p
	res[2] = p + 1

	if size%2 == 1 {
		p /= 10
	}

	for p > 0 {
		res[0] = res[0]*10 + p%10
		res[1] = res[1]*10 + p%10
		res[2] = res[2]*10 + p%10
		p /= 10
	}

	return res
}

func getBitChangeCandidates(size int) []int {
	base := 1
	for i := 1; i < size; i++ {
		base *= 10
	}

	res := make([]int, 4)
	res[0] = base - 1
	res[1] = base + 1

	base *= 10

	res[2] = base - 1
	res[3] = base + 1

	return res
}
