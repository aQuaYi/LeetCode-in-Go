package Problem0759

import (
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	// 这样排序的意义是
	// 假设 intervals[i:j] 中的 b 都是一样的
	// [b-1,b] 和 intervals[i:j] 中的每一个也肯定都有两个公共元素
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	// left 和 right 取过的值，就是 set 中的值
	left := intervals[0][1] - 1
	right := intervals[0][1]
	res += 2

	n := len(intervals)

	for i := 1; i < n; i++ {
		curr := intervals[i]
		if left < curr[0] && curr[0] <= right {
			res++
			// 原先的 left 没有包含在 curr 中
			// 新 left = right
			left = right
			// 新 right = curr[1]
			right = curr[1]
		} else if curr[0] > right {
			// left 和 right 都没有包含在 curr 中
			res += 2
			left = curr[1] - 1
			right = curr[1]
		}
	}

	return res
}
