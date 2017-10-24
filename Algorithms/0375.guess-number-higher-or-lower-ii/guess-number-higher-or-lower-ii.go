package Problem0375

func getMoneyAmount(n int) int {

	// dp[i][j] 保证能猜出 i<=x<=j 中 x 的具体值的最小金额
	// dp[1][n] 是答案
	dp := make([]int, n+1)

	MIN := 1<<63 - 1

	var i, j, k int
	for j = 2; j <= n; j++ {
		for i = 1; i < j; i++ {
			dp[j] = MIN
			for k = 1; k < i; k++ {
				dp[j] = min(dp[j], k+max(dp[k-1], k+dp[i-k]))
			}
		}
	}

	return dp[n]
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
