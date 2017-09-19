package Problem0128

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)

	max, temp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			temp++
		} else if max < temp {
			max = temp
		}
	}

	if temp == len(nums) {
		return len(nums)
	}

	return max
}
