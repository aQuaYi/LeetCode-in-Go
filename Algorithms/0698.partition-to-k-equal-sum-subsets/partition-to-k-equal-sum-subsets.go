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

	if sum%k != 0 || sum/k < max {
		// 提前结束
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// isUsed[i] == true 表示 nums[i] 已经被归类到某个小组，无法再使用了
	isUsed := make([]bool, size)

	return dfs(nums, isUsed, 0, k, 0, 0, sum/k)
}

func dfs(nums []int, isUsed []bool, startIndex, k, curSum, curNum, target int) bool {
	if k == 1 {
		return true
	}

	if curSum == target && curNum > 0 {
		return dfs(nums, isUsed, 0, k-1, 0, 0, target)
	}

	for i := startIndex; i < len(nums); i++ {
		if !isUsed[i] && curSum+nums[i] <= target {
			isUsed[i] = true
			if dfs(nums, isUsed, i+1, k, curSum+nums[i], curNum+1, target) {
				return true
			}
			isUsed[i] = false
		}
	}

	return false
}
