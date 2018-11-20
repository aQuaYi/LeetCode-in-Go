package problem0910

import (
	"sort"
)

// ref: https://leetcode.com/problems/smallest-range-ii/discuss/173377/C++JavaPython-Add-0-or-2-*-K

func smallestRangeII(A []int, k int) int {
	sort.Ints(A)

	n := len(A)
	mx, mn := A[n-1], A[0]
	res := mx - mn

	for i := 0; i+1 < n; i++ {
		mx = max(mx, A[i]+2*k)
		mn = min(A[i+1], A[0]+2*k)
		res = min(res, mx-mn)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
