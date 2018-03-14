package problem0611

import (
	"sort"
)

func triangleNumber(a []int) int {
	sort.Ints(a)
	n := len(a)

	res := 0
	for k := n - 1; 2 <= k; k-- {
		for i, j := 0, k-1; i < j; {
			// 由于 a 是递增的，保证了 a[i] <= a[j] <= a[k]
			// 所以，只要 a[i] + a[j] > a[k]
			// 三者就可以构成三角形
			if a[i]+a[j] > a[k] {
				// a[j], a[k] 和 a[i:j] 中的任意一个，都可以组成三角形
				// a[i:j] 中元素个数是 j-i
				res += j - i
				// a[j] 足够长，可以变短试试看
				j--
			} else {
				// a[i] 不够长，需要变长后，再看看
				i++
			}
		}
	}

	return res
}
