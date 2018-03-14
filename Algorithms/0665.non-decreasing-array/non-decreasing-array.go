package problem0665

import (
	"sort"
)

func checkPossibility(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			pre := deepCopy(nums)
			// 把 nums[i-1] 修改成　nums[i]
			pre[i-1] = pre[i]
			next := deepCopy(nums)
			// 把 nums[i] 修改成　nums[i-1]
			next[i] = next[i-1]
			// 修改完成后，查看修改完成后，能否满足题意
			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}

	return true
}

func deepCopy(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}
