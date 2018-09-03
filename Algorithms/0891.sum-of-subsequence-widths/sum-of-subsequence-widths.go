package problem0891

import (
	"sort"
)

const mod = 1e9 + 7

func sumSubseqWidths(a []int) int {
	sort.Ints(a)
	res := 0

	n := len(a)
	c := 1
	for i := 0; i < n; i++ {
		res = (res + a[i]*c - a[n-i-1]*c) % mod
		c = (c << 1) % mod
	}

	return res % mod
}
