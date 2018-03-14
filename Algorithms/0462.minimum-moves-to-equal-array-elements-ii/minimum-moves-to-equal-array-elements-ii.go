package problem0462

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	size := len(nums)

	median := nums[size/2]
	if size&1 == 0 {
		median += (nums[size/2-1] - median) / 2
	}

	res := 0
	for _, n := range nums {
		// 把所有的元素都调整到 中位数 才是最优解
		res += diff(median, n)
	}

	return res
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
