package Problem0324

import "sort"

func wiggleSort(nums []int) {
	sort.Ints(nums)

	offset := (len(nums) + 1) / 2
	if offset%2 != 1 {
		offset++
	}
	i := 1
	for i < len(nums)/2 {
		nums[i], nums[i+offset] = nums[i+offset], nums[i]
		i += 2
	}

	return
}
