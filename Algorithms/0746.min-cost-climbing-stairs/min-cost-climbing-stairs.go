package problem0746

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// dp[i] 表示，爬到 i-th 台阶需要的花费
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
