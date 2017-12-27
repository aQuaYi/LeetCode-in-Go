package Problem0719

import "sort"

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
		if count(a, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func count(a []int, mid int) int {
	n := len(a)
	res := 0
	i, j := 0, 1
	for j < n {
		if a[j]-a[i] <= mid {
			res += j - i
			j++
		} else {
			i++
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
