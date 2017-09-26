package Problem0164

import "sort"

func maximumGap(nums []int) int {
	sort.Ints(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		temp := nums[i] - nums[i-1]
		if max < temp {
			max = temp
		}
	}

	return max
}
