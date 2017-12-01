package Problem0664

func strangePrinter(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	// dp[i][j] 表示打印 s[i:j+1] 所需的最小次数
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for Len := 2; Len <= n; Len++ {
		for i := 0; i+Len-1 < n; i++ {
			j := i + Len - 1
			dp[i][j] = Len
			for k := i + 1; k <= j; k++ {
				temp := dp[i][k-1] + dp[k][j]
				if s[k-1] == s[j] {
					temp--
				}
				dp[i][j] = min(dp[i][j], temp)
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
