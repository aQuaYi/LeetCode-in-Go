package Problem0664

func strangePrinter(s string) int {
	n := len(s)
	if n < 1 {
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
			//
			left := 1
			if s[i] == s[i+1] || s[i] == s[j] {
				left = 0
			}
			//
			right := 1
			if s[j-1] == s[j] || s[j] == s[i] {
				right = 0
			}
			//
			dp[i][j] = min(left+dp[i+1][j], dp[i][j-1]+right)
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
