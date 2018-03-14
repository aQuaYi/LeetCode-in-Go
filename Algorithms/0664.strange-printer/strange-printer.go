package problem0664

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
			// dp[i][j] 可以由 dp[i][j-1] 打印而来
			dp[i][j] = dp[i][j-1] + 1
			if s[j-1] == s[j] {
				dp[i][j]--
			}
			//
			for k := i + 1; k <= j; k++ {
				// 当 s[k-1] == s[j] 时，可以一次打印 s[k-1:j+1] 再修改 s[k:j] 上的内容
				// 这样也可以少打印一次
				if s[k-1] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k][j]-1)
				}
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
