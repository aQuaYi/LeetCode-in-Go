package Problem0041

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {

		fmt.Println(i)

		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			//  0 <= nums[i]-1 && nums[i]-1 < len(nums) 是为了防止nums[nums[i]-1]溢出
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
