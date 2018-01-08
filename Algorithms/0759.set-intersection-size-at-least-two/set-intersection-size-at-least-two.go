package Problem0759

import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	hasNum := make(map[int]bool, n*2)

	interSect := intervals[0]
	for i := 1; i < n; i++ {
		if interSect[1] <= intervals[i][0] {
			hasNum[interSect[0]] = true
			hasNum[interSect[1]] = true
			interSect = intervals[i]
			continue
		}

		interSect[0] = intervals[i][0]
		if intervals[i][1] < interSect[1] {
			interSect[1] = intervals[i][1]
		}
	}
	hasNum[interSect[0]] = true
	hasNum[interSect[1]] = true

	return len(hasNum)
}
