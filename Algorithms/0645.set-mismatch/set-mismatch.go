package Problem0645

import "sort"

func findErrorNums(nums []int) []int {
	sort.Ints(nums)

	res := make([]int, 2)
	res[1] = len(nums)

	pre := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			res[0] = pre
		} else if nums[i] == pre+2 {
			res[1] = pre + 1
		}

		pre = nums[i]
	}

	return res
}
