package Problem0665

import (
	"sort"
)

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			nums[i] = nums[i-1]
			return sort.IsSorted(sort.IntSlice(nums))
		}
	}

	return true
}
