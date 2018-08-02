package problem0870

import (
	"sort"
)

func advantageCount(A []int, B []int) []int {
	size := len(A)

	BI := make([][2]int, size)
	for i, n := range B {
		BI[i][0], BI[i][1] = n, i
	}

	sort.Slice(BI, func(i int, j int) bool {
		return BI[i][0] < BI[j][0]
	})

	res := make([]int, size)

	sort.Ints(A)

	l, r := 0, size-1
	for _, a := range A {
		if BI[l][0] < a {
			res[BI[l][1]] = a
			l++
		} else {
			res[BI[r][1]] = a
			r--
		}
	}

	return res
}
