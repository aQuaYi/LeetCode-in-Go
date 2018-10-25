package problem0078

import (
	"math"
)

func subsets(nums []int) [][]int {
	size := resSize(nums)
	res := make([][]int, 0, size)
	subs := make(chan []int, 16)
	go rec(nums, []int{}, subs)
	for sub := range subs {
		res = append(res, sub)
		size--
		if size == 0 {
			close(subs)
		}
	}
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, c chan []int) {
	size := len(nums)
	if size == 0 {
		c <- temp
		return
	}

	// 对于每个元素来说，要么存在于某个子集中，要么不存在

	// 把 nums 中的最后一个元素，不放入 temp
	go rec(nums[:size-1], temp, c)
	// 把 nums 中的最后一个元素，*放入* temp
	newTemp := make([]int, 1, len(nums)+len(temp))
	newTemp[0] = nums[size-1]
	newTemp = append(newTemp, temp...)
	go rec(nums[:size-1], newTemp, c)
}
