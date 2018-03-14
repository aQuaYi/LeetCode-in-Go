package problem0274

import "sort"

func hIndex(d []int) int {
	// 对 citations 进行降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	size := len(d)

	// 二分查找法
	lo, hi := 0, size-1
	var mi int
	for lo <= hi {
		mi = (lo + hi) / 2
		if d[mi] > mi {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return lo
}
