package Problem0274

import "sort"

func hIndex(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	size := len(citations)
	h := 0
	for h < size && citations[h] > h {
		h++
	}

	return h
}
