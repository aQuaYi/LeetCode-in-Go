package problem0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	// dp[i][j][i+j] == true 表示 s1[:i] 和 s2[:j] 可以生成 s3[:i+j]
	dp := make([][][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]bool, m+n+1)
		}
	}
	dp[0][0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0][i] = dp[i-1][0][i-1] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= n; j++ {
		dp[0][j][j] = dp[0][j-1][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j][i+j] = (dp[i-1][j][i+j-1] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1][i+j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n][m+n]
}
