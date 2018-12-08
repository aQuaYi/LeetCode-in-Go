package problem0904

// 找到最长的 tree 的子数组，要求最多只能含有两种数字
func totalFruit(trees []int) int {
	// 子数组 "252555" 中， tail 是尾部的 "555"
	// tailType 是 “5”
	// tailCount 是 3
	tailType, theOther := -1, -1
	tailCount, count := 0, 0 // 子数组尾部相同类型水果的长度，子数组的长度
	res := 0

	for _, t := range trees {
		count++
		if t != tailType && t != theOther {
			// t 是新的类型
			// 原先的子数组的 tail 加上 t 组成新的子数组
			count = tailCount + 1
		}

		res = max(res, count)

		tailCount++
		if t != tailType {
			// t 成为新的 tail，需要更新 tailType 和 tailCount
			theOther, tailType = tailType, t
			tailCount = 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ref: https://leetcode.com/problems/fruit-into-baskets/discuss/170745/Problem:-Longest-Subarray-With-2-Elements
