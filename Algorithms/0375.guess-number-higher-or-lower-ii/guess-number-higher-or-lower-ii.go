package Problem0375

func getMoneyAmount(n int) int {

	// dp[i] 保证能猜出 1<=x<=i 中 x 的具体值的最小金额
	// dp[n] 是答案
	dp := make([]int, n+1)
	var i, j int
	for i = 2; i <= n; i++ {
		dp[i] = i - 1 + dp[i-2]
		for j = i - 2; 1 < j; j-- {
			dp[i] = min(dp[i], j+max(dp[j-1], j+dp[i-j]))
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
