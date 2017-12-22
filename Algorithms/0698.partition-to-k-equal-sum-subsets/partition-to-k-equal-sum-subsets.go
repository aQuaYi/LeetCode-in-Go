package Problem0698

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	size := len(nums)
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if max < n {
			max = n
		}
	}

	if sum%4 != 0 || sum/size < max {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	isUsed := make([]bool, size)
	return dfs(nums, isUsed, 0, 0, 0, 0, sum/size)
}

func dfs(nums []int, isUsed []bool, startIndex, k, curSum, curNum, target int) bool {
	if k == 1 {
		return true
	}

	if curSum == target && curNum > 0 {
		return dfs(nums, isUsed, 0, k-1, 0, 0, target)
	}

	for i := startIndex; i < len(nums); i++ {
		if !isUsed[i] {
			isUsed[i] = true
			if dfs(nums, isUsed, i+1, k, curSum+nums[i], curNum+1, target) {
				return true
			}
			isUsed[i] = false
		}

	}

	return false
}
