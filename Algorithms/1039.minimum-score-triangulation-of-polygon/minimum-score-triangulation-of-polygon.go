package problem1039

func minScoreTriangulation(A []int) int {
	n := len(A)
	A = append(A, A[0])
	dp := [51][51]int{}

	for r := 2; r <= n; r++ {
		for i := 1; i+r-1 <= n; i++ {
			j := i + r - 1
			dp[i][j] = 1<<63 - 1
			for k := i; k <= j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+A[i-1]*A[k]*A[j])
			}
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
