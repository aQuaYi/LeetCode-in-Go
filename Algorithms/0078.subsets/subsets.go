package problem0078

import (
	"sort"
)

func subsets(nums []int) [][]int {
	res := [][]int{}

	recur(nums, []int{}, &res)

	return res
}

func recur(nums, temp []int, res *[][]int) {
	l := len(nums)
	if l == 0 {
		sort.Ints(temp)
		*res = append(*res, temp)
		return
	}

	recur(nums[:l-1], temp, res)

	recur(nums[:l-1], append([]int{nums[l-1]}, temp...), res)
}
