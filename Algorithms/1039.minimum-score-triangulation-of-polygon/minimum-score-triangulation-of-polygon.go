package problem1039

func minScoreTriangulation(A []int) int {
	n := len(A)
	dp := [50][50]int{}
	for d := 2; d < n; d++ {
		for i := 0; i+d < n; i++ {
			k := i + d
			dp[i][k] = 1 << 30
			for j := i + 1; j < k; j++ {
				dp[i][k] = min(dp[i][k], dp[i][j]+dp[j][k]+A[i]*A[j]*A[k])
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
