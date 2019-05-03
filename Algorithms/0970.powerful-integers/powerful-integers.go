package problem0970

import (
	"sort"
)

func powerfulIntegers(x int, y int, bound int) []int {
	if x == 1 {
		x = bound + 1
	}
	if y == 1 {
		y = bound + 1
	}

	res := make([]int, 0, 128)
	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			res = append(res, i+j)
		}
	}

	return removeRepeated(res)
}

func removeRepeated(nums []int) []int {
	sort.Ints(nums)

	size := len(nums)

	last, j := -1, -1
	for i := 0; i < size; i++ {
		if last == nums[i] { // nums[i]>0 for any i
			continue
		}
		j++
		nums[j], last = nums[i], nums[i]
	}
	return nums[:j+1]
}
