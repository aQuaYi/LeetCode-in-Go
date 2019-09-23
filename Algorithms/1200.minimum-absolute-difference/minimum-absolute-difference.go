package problem1200

import (
	"math"
	"sort"
)

func minimumAbsDifference(A []int) [][]int {
	sort.Ints(A)
	n := len(A)

	minDiff := math.MaxInt64
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, A[i]-A[i-1])
	}

	res := make([][]int, 0, n)

	for i := 1; i < n; i++ {
		if A[i]-A[i-1] == minDiff {
			res = append(res, []int{A[i-1], A[i]})
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
