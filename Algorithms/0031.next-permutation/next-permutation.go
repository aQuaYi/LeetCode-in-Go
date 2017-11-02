package Problem0031

import "sort"

func nextPermutation(nums []int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	var i int
	for i = length - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}

	if i > 0 {
		sort.Ints(nums[i:])
		for j := i - 1; j < length; j++ {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				return
			}
		}
	}

	sort.Ints(nums)

}
