package problem0718

func findLength(a1, a2 []int) int {
	n1, n2 := len(a1), len(a2)

	// dp[i][j] == k  表示 a1[i-k:i] == a2[j-k:j] ，
	// 但是 a1[i-k-1] != a2[j-k-1]
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	res := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if a1[i-1] == a2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
