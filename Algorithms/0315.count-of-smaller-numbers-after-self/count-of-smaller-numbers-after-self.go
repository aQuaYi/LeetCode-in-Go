package Problem0315

import (
	"sort"
)

func countSmaller(nums []int) []int {
	size := len(nums)
	temp := make([]int, size)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		temp = temp[:size-i-1]
		copy(temp, nums[i+1:])
		sort.Ints(temp)

		res[i] = sort.SearchInts(temp, nums[i])
	}

	return res
}
