package Problem0274

import "sort"

func hIndex(citations []int) int {
	// 对 citations 进行降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	size := len(citations)
	h := 0
	// h 的含义是
	// 前 h 个的 citation 都 > h
	for h < size && citations[h] > h {
		h++
	}

	return h
}
