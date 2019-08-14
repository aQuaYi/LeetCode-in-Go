package problem1155

const mod = 1e9 + 7

func numRollsToTarget(d int, f int, target int) int {
	dp := [31][1001]int{}
	dp[0][0] = 1

	for i := 1; i <= d; i++ {
		maxJ := min(target, i*f)
		for j := i; j <= maxJ; j++ {
			maxK := min(f, j)
			for k := 1; k <= maxK; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
			dp[i][j] %= mod
		}
	}

	return dp[d][target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
