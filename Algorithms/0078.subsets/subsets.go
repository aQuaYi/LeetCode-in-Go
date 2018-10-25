package problem0078

import (
	"math"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0, resSize(nums))
	rec(nums, []int{}, &res)
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	if size == 0 {
		*res = append(*res, temp)
		return
	}

	// 对于每个元素来说，要么存在于某个子集中，要么不存在

	// 把 nums 中的最后一个元素，不放入 temp
	rec(nums[:size-1], temp, res)
	// 把 nums 中的最后一个元素，*放入* temp
	rec(nums[:size-1], append([]int{nums[size-1]}, temp...), res)
}
