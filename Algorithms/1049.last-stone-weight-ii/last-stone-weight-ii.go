package problem1049

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	dp := [32][32]int{}

	for width := 1; width <= n; width++ {
		for i := 1; i+width-1 <= n; i++ {
			j := i + width - 1
			dp[i][j] = min(abs(stones[i-1]-dp[i+1][j]),
				abs(stones[j-1]-dp[i][j-1]))
		}
	}

	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
