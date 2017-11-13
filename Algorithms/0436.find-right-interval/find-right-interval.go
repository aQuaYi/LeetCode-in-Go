package Problem0436

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"sort"
)

type Interval = kit.Interval

func findRightInterval(intervals []Interval) []int {
	if len(intervals) == 0 {
		return []int{}
	}

	starts := make([]int, len(intervals))
	indexes := make(map[int]int, len(intervals))
	result := make([]int, len(intervals))

	for i, v := range intervals {
		starts[i] = v.Start
		indexes[v.Start] = i
	}

	sort.Ints(starts)

	for i, v := range intervals {
		idx := sort.SearchInts(starts, v.End)
		if idx < len(intervals) {
			result[i] = indexes[starts[idx]]
		} else {
			result[i] = -1
		}
	}

	return result
}
