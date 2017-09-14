package Problem0115

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	// dp[i][j] == numDistinct(s[:i], t[:j])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		// numDistinct(s[:i], "") == 1
		// s[:i] 删除所有的字符后，才会变成 ""，只有这一种可能。
		dp[i][0] = 1
	}

	for j := 1; j <= n; j++ {
		for i := j; i <= m; i++ {
			// 对于 dp[i][j] 来说
			// s[:i] 中符合条件的子字符串，按照是否包含 s[i-1]，可以分成两个部分：
			// 第一部分：
			//     **不**包含 s[i-1]，
			//     这部分的数量，等于 dp[i-1]dp[j]
			dp[i][j] = dp[i-1][j]
			// 第二部分：
			//     包含 s[i-1]
			//     这部分，只有当 s[i-1] == t[j-1] 的时候，才存在
			//     存在的话，这部分的数量，等于 dp[i-1][j-1]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
