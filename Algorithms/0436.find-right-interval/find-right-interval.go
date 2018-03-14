package problem0436

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"sort"
)

type Interval = kit.Interval

func findRightInterval(intervals []Interval) []int {
	size := len(intervals)

	starts := make([]int, size)
	idxs := make(map[int]int, size)
	res := make([]int, size)

	for i, v := range intervals {
		starts[i] = v.Start
		idxs[v.Start] = i
	}

	sort.Ints(starts)

	for i, v := range intervals {
		idx := sort.SearchInts(starts, v.End)
		if idx < size {
			res[i] = idxs[starts[idx]]
		} else {
			res[i] = -1
		}
	}

	return res
}
