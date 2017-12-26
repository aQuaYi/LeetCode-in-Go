package Problem0712

func minimumDeleteSum(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)

	// dp[i][j] 表示 s1[:i] 和 s2[:j] 的最大 ASCII 公共子序列
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			}
		}
	}

	return sum(s1) + sum(s2) - dp[n1][n2]*2
}

// 返回 s 中所在字母的 ASCII 码之和
func sum(s string) int {
	res := 0
	for i := range s {
		res += int(s[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
