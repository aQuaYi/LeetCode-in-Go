package Problem0581

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	i, j := 0, len(nums)-1

	for i < len(nums) && temp[i] == nums[i] {
		i++
	}

	for i < j && temp[j] == nums[j] {
		j--
	}

	return j - i + 1
}
