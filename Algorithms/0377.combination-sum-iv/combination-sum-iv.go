package problem0377

func combinationSum4(nums []int, target int) int {
	// dp[i] 是 i 能用 nums 中的数所组合的种类数
	// dp[target] 是答案
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := range nums {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
