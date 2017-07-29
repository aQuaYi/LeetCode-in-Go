package Problem0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			l, r := j+1, len(nums)-1
			for l < r {

				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s < target:
					l++
				case s > target:
					r--
				default:
				}
			}
		}

	}
	return res
}
