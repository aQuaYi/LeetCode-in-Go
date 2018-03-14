package problem0698

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

	// nums 降序排列可以加快 dfs 的速度
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// isUsed[i] == true 表示 nums[i] 已经被归类到某个小组，无法再使用了
	isUsed := make([]bool, size)

	return dfs(nums, isUsed, k, 0, 0, sum/k)
}

func dfs(nums []int, isUsed []bool, k, start, sum, target int) bool {
	if k == 1 {
		// 已经找到了 k-1 组解
		// 剩下的自然就是第 k 组解
		return true
	}

	if sum == target {
		// 找到第 k 组的一种解
		// 开始搜索 k-1 组的解
		return dfs(nums, isUsed, k-1, 0, 0, target)
	}

	for i := start; i < len(nums); i++ {
		if !isUsed[i] && sum+nums[i] <= target {
			isUsed[i] = true
			// 试着 sum+nums[i]
			if dfs(nums, isUsed, k, i+1, sum+nums[i], target) {
				return true
			}
			isUsed[i] = false
		}
	}

	return false
}
