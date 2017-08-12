package Problem0041

import "sort"

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	sort.Ints(nums)
	i := sort.SearchInts(nums, 0)
	if nums[i] == 0 {
		i++
	}

	res := 1
	for i < len(nums) {
		if 0 < i && nums[i] == nums[i-1] {
			i++
			continue
		}
		if nums[i] != res {
			return res
		}
		i++
		res++
	}

	return res
}
