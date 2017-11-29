package Problem0486

// PredictTheWinner 在 play1 能赢的时候，返回 true
func PredictTheWinner(nums []int) bool {
	n := len(nums)

	// dp[i][j] 表示 nums[i:j+1] 中 play1 比 play2 多的得分
	// dp[0][n-1] >=0 表示 play1 获胜
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for Len := 1; Len < n; Len++ {
		for left := 0; left < n-Len; left++ {
			right := left + Len
			dp[left][right] = max(nums[left]-dp[left+1][right], nums[right]-dp[left][right-1])
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
