package problem1200

import (
	"math"
	"sort"
)

func minimumAbsDifference(A []int) [][]int {
	sort.Ints(A)
	n := len(A)
	isExist := make(map[int]bool, n)

	minDiff := math.MaxInt64
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, A[i]-A[i-1])
		isExist[A[i]] = true
	}

	res := make([][]int, 0, n)

	for i := 0; i < n-1; i++ {
		if !isExist[A[i]+minDiff] {
			continue
		}
		res = append(res, []int{A[i], A[i] + minDiff})
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
