package problem0516

func longestPalindromeSubseq(s string) int {
	n := len(s)

	// dp[i][j] 表示 s[i:j+1] 中最长的回文子字符串
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for Len := 2; Len <= n; Len++ {
		for i := 0; i+Len-1 < n; i++ {
			j := i + Len - 1
			if s[i] == s[j] {
				// 首位相等，则在原先的长度长 +2
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 首位不等，则在等于去除首尾后的较长者
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
