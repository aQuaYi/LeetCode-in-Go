package problem0486

// PredictTheWinner 在 play1 能赢的时候，返回 true
func PredictTheWinner(nums []int) bool {
	n := len(nums)

	// dp[i][j] 表示 nums[i:j+1] 中 play1 比 play2 多的得分
	// dp[0][n-1] >=0 表示 play1 获胜
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		// 只有 nums[i] 时，play1 比 play2 多 nums[i] 分
		// 因为 play1 选了 nums[i]，而 play2 无数可选了
		dp[i][i] = nums[i]
	}

	for Len := 2; Len <= n; Len++ {
		for i := 0; i <= n-Len; i++ {
			// Len == j - i + 1
			j := i + Len - 1

			// dp[i][j] 表示 nums[i:j+1] 中 play1 比 play2 多的得分
			// 那么 -dp[i][j] 表示什么呢?
			// -dp[i][j] 表示 nums[i:j+1] 中 play2 比 play1 多的得分
			// 所以，
			// nums[i]-dp[i+1][j] 等于 play1 选择 左端 的数后比 play2 多的得分
			// nums[j]-dp[i][j-1] 等于 play1 选择 右端 的数后比 play2 多的得分
			// 从以上两者之间，选择一个大的，作为 dp[i][j]
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
