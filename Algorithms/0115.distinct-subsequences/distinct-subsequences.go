package Problem0115

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	if n == 0 {
		return 1
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}

	for j := 1; j <= n; j++ {
		if s[:j] == t[:j] {
			dp[j][j] = 1
		}

		for i := j + 1; i <= m; i++ {
			if n-j <= m-i && s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}
	return dp[m][n]
}
