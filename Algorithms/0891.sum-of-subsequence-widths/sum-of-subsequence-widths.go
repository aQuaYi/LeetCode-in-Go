package problem0891

import (
	"sort"
)

const mod = 1e9 + 7

func sumSubseqWidths(a []int) int {
	sort.Ints(a)
	res := 0

	size := len(a)
	for i := 0; i < size; i++ {
		base := 1
		for j := i + 1; j < size; j++ {
			res += (a[j] - a[i]) * base
			base *= 2
		}
		res %= mod
	}
	return res
}
