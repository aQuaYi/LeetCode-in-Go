package Problem0759

import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	left := intervals[0][1] - 1
	right := intervals[0][1]
	res += 2

	n := len(intervals)

	for i := 1; i < n; i++ {
		curr := intervals[i]
		if left < curr[0] && curr[0] <= right {
			res++
			left = right
			right = curr[1]
		} else if curr[0] > right {
			res += 2
			left = curr[1] - 1
			right = curr[1]
		}
	}

	return res
}
