package problem1155

const mod = 1e9 + 7

func numRollsToTarget(d int, f int, target int) int {
	dp := [31][1001]int{}
	dp[0][0] = 1

	for i := 1; i <= d; i++ {
		max := i * f
		for j := i; j <= target && j <= max; j++ {
			for k := 1; k <= f && k <= j; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
			dp[i][j] %= mod
		}
	}

	return dp[d][target]
}
