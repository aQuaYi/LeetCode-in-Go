package Problem0719

import (
	"sort"
)

func smallestDistancePair(a []int, k int) int {
	n := len(a)
	sort.Ints(a)

	low := a[1] - a[0]
	for i := 2; i < n; i++ {
		low = min(low, a[i]-a[i-1])
	}

	high := a[n-1] - a[0]

	for low < high {
		mid := low + (high-low)/2
		if countPairs(a, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func countPairs(a []int, mid int) int {
	n := len(a)
	res := 0
	for i := 0; i < n; i++ {
		j := i
		for j < n && a[j]-a[i] <= mid {
			j++
		}
		res += j - i - 1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
