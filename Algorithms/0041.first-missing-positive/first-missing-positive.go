package Problem0041

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := nums[i] - 1
		if nums[i] < 1 || nums[i] > len(nums) || index < 0 || index > len(nums) || nums[i] == i+1 {
			// cannot swap or noused
			continue
		}
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			fmt.Println(nums)
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
