package problem0473

import "sort"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}

	// 先对nums中的数做一个基本的分析
	ok, avg := analyze(nums)
	if !ok {
		return false
	}

	// 降序排列可以加快 dfs 的速度
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	edges := make([]int, 4)
	return dfs(nums, edges, 0, avg)
}

func dfs(nums, edges []int, index, target int) bool {
	if index == len(nums) {
		if edges[0] == target &&
			edges[1] == target &&
			edges[2] == target {
			return true
		}
		// TODO: 重新组织程序，达到完整的覆盖率
		return false
	}

	for i := 0; i < 4; i++ {
		if edges[i]+nums[index] > target {
			continue
		}
		edges[i] += nums[index]
		if dfs(nums, edges, index+1, target) {
			return true
		}
		edges[i] -= nums[index]
	}

	return false
}

// analyze 统计了 nums 中数的总和，能否被 4 整除
func analyze(nums []int) (ok bool, avg int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 {
		return false, 0
	}
	return true, sum / 4
}
