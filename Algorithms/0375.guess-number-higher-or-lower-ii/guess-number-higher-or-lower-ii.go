package problem0375

func getMoneyAmount(n int) int {
	// dp[i][j] 保证能猜出 i<=x<=j 中 x 的具体值的最小金额
	// dp[1][n] 是答案
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for j := 2; j <= n; j++ {
		for i := j - 1; 0 < i; i-- {
			// 为了确保可以猜出 i<=x<=j 中的 x
			// 第一次，我们可以猜 x 为，i，i+1，...,j-1
			// 所有这些可能性中的最小值就是 dp[i][j] 的值
			dp[i][j] = i + dp[i+1][j]
			for k := i + 1; k < j; k++ {
				// k+max(dp[i][k-1], dp[k+1][j])) 猜 x 为 k 所花费的最小费用
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
