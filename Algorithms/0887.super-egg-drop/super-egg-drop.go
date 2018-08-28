package problem0887

func superEggDrop(K, N int) int {
	dp := [101]int{}
	step := 0
	for dp[K] < N {
		for k := K; k > 0; k-- {
			dp[k] = dp[k-1] + dp[k] + 1
		}
		step++
	}
	return step
}
