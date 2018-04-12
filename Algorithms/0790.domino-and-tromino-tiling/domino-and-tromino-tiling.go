package problem0790

const mod = 1e9 + 7

func numTilings(N int) int {
	dp := [1001]int{1: 1, 2: 2, 3: 5}
	if N <= 3 {
		return dp[N]
	}

	for i := 4; i <= N; i++ {
		dp[i] = 2*dp[i-1] + dp[i-3]
		dp[i] %= mod
	}

	return dp[N]
}
