package Problem0260

import (
	"sort"
)

func singleNumber(nums []int) []int {
	sort.Ints(nums)
	res := make([]int, 2)
	idx := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			i++
			continue
		} else {
			res[idx] = nums[i]
			idx++
			if idx == 2 {
				return res
			}
		}
	}

	if idx == 1 {
		res[1] = nums[len(nums)-1]
	}

	return res
}
