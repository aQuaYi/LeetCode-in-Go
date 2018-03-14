package problem0128

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
		} else if nums[i] != nums[i-1] {
			// [0,1,1,2]　仍然可以视为连续
			temp = 1
		}
		//　更新　max
		if max < temp {
			max = temp
		}
	}

	return max
}
