package Problem0494

func findTargetSumWays(nums []int, S int) int {
	res := 0
	dfs(nums, 0, 0, S, &res)
	return res
}

func dfs(nums []int, index, sum, target int, res *int) {
	if index == len(nums) {
		if sum == target {
			*res++
		}
		return
	}

	dfs(nums, index+1, sum+nums[index], target, res)
	dfs(nums, index+1, sum-nums[index], target, res)
}
