package Problem0078

import (
	"sort"
)

func subsets(nums []int) [][]int {
	res := [][]int{}

	cur(nums, []int{}, &res)

	return res
}

func cur(nums, temp []int, res *[][]int) {
	l := len(nums)
	if l == 0 {
		sort.Ints(temp)
		*res = append(*res, temp)
		return
	}

	cur(nums[:l-1], temp, res)

	cur(nums[:l-1], append([]int{nums[l-1]}, temp...), res)
}
