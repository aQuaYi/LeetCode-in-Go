package problem0518

func change(amount int, coins []int) int {
	size := len(coins)
	// dp[i][j] 表示，使用 coins[:i] 中的零钱，组成金额为 j 的总种类数
	dp := make([][]int, size+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1

	for i := 1; i <= size; i++ {
		dp[i][0] = 1
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j] // 不使用 coins[i-1] 时候的种类数
			if j >= coins[i-1] {
				// 此外，如果
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[size][amount]
}
