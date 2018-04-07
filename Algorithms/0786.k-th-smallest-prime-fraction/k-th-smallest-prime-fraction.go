package problem0786

import (
	"sort"
)

func kthSmallestPrimeFraction(A []int, K int) []int {
	n := len(A)

	if n == 2 {
		return A
	}

	sort.Ints(A)

	fracs := make([][]int, 0, n)

	for i := n - 1; i > 0; i-- {
		for j := 0; j < n-1; j++ {
			fracs = append(fracs, []int{A[j], A[i]})
		}
	}

	sort.Slice(fracs, func(i int, j int) bool {
		return isLess(fracs[i], fracs[j])
	})

	return fracs[K-1]
}

// 如果 a < b ，则返回 true
func isLess(a, b []int) bool {
	return a[0]*b[1] < a[1]*b[0]
}
