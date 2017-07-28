package Problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					if yes := contain(res, temp); !yes {
						res = append(res, temp)

					}
				}
			}
		}
	}

	return res
}

func contain(numss [][]int, nums []int) bool {
	sort.Ints(nums)

	for _, ns := range numss {
		sort.Ints(ns)
		if equal(ns, nums) {
			return true
		}
	}

	return false
}

func equal(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
