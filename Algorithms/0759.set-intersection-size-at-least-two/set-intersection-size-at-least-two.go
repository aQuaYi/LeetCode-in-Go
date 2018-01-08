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
	lastA := intervals[n-1][0]

	hasNum := make(map[int]bool, len(intervals)*2)
	hasNum[lastA] = true
	hasNum[lastA+1] = true

	for i := 0; i < n-1; i++ {
		update(hasNum, intervals[i])
	}

	return len(hasNum)
}

func update(hasNum map[int]bool, interval []int) {
	countDown := 2
	a, b := interval[0], interval[1]
	for i := a; i <= b && countDown > 0; i++ {
		if hasNum[i] {
			countDown--
		} else {
			break
		}
	}

	for countDown > 0 {
		hasNum[b] = true
		b--
		countDown--
	}

	return
}
