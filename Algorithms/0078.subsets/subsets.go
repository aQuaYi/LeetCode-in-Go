package problem0078

import (
	"sort"
)

func subsets(nums []int) [][]int {
	res := [][]int{}
	rec(nums, []int{}, &res)
	return res
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	if size == 0 {
		sort.Ints(temp)
		*res = append(*res, temp)
		return
	}

	rec(nums[:size-1], temp, res)

	rec(nums[:size-1], append([]int{nums[size-1]}, temp...), res)
}
