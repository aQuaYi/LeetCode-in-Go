package problem0870

import (
	"sort"
)

func advantageCount(A []int, B []int) []int {
	size := len(A)
	isTaken := make([]bool, size)

	search := func(n int) int {
		l, r := 0, size-1

		for l < r {
			m := (l + r) / 2
			if A[m] <= n {
				l = m + 1
			} else {
				r = m
			}
		}

		for isTaken[r] {
			r = (r + 1) % size
		}

		isTaken[r] = true

		return r
	}

	sort.Ints(A)

	res := make([]int, size)

	for i, n := range B {
		res[i] = A[search(n)]
	}

	return res
}
