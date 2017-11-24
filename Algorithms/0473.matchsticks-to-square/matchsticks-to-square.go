package Problem0473

import "sort"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}

	ok, avg := isOK(nums)
	if !ok {
		return false
	}

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

func isOK(nums []int) (ok bool, avg int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 {
		return false, 0
	}
	return true, sum / 4
}
