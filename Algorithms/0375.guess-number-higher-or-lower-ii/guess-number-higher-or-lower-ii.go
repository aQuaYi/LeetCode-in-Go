package Problem0375

func getMoneyAmount(n int) int {

	// dp[i][j] 保证能猜出 i<=x<=j 中 x 的具体值的最小金额
	// dp[1][n] 是答案
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	MIN := 1<<63 - 1

	var i, j, k int
	for j = 2; j <= n; j++ {
		for i = j - 1; 0 < i; i-- {
			Min := MIN
			for k = i + 1; k < j; k++ {
				Min = min(Min, k+max(dp[i][k-1], dp[k+1][j]))
			}
			if i+1 == j {
				dp[i][j] = i
			} else {
				dp[i][j] = Min
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
