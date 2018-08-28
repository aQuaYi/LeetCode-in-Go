package problem0887

func superEggDrop(K, N int) int {
	dp := [10001][101]int{}
	m := 0
	for dp[m][K] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[m][k] = dp[m-1][k-1] + dp[m-1][k] + 1
		}
	}
	return m
}
