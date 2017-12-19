package Problem0658

import (
	"sort"
)

func findClosestElements(a []int, k int, x int) []int {
	i := sort.SearchInts(a, x)
	if i == 0 && x < a[0] {
		return a[:k]
	}
	if i >= len(a)-1 {
		return a[len(a)-k:]
	}

	l, r := i-1, i
	for r-l-1 < k {
		if l == -1 {
			r++
			continue
		}

		if r == len(a) {
			l--
			continue
		}

		if a[l]+a[r] >= 2*x {
			l--
		} else {
			r++
		}
	}

	return a[l+1 : r]
}
