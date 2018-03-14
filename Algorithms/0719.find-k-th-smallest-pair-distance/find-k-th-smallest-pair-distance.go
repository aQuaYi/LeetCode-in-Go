package problem0719

import "sort"

func smallestDistancePair(a []int, k int) int {
	n := len(a)
	sort.Ints(a)

	low := a[1] - a[0]
	for i := 2; i < n; i++ {
		low = min(low, a[i]-a[i-1])
	}

	high := a[n-1] - a[0]
	// 由题意可知，res 必定存在于 [low, high] 之中
	// 又由 count(a, mid) 可以 a 中距离 <= mid 的个数
	// 故可以使用 二分查找法
	for low < high {
		mid := low + (high-low)/2
		if count(a, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
		// NOTICE:
		// count(a,mid) == k , 并不能立即返回 mid 作为答案
		// 例如 count([]int{1,1,4}, 2) == 1,
		// 但是 smallestDistancePair([]int{1,1,4}, 1) == 0
	}

	return low
}

// 统计 a 中距离 <= mid 的元素对的数量
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
