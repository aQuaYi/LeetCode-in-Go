package problem0044

func isMatch(s string, p string) bool {
	sSize, pSize := len(s), len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}
	// dp[i][j] == true 意味着，s[:i+1] 可以和 p[:j+1] 匹配
	dp[0][0] = true

	for j := 1; j <= pSize; j++ {
		if p[j-1] == '*' {
			// 当 p[j-1] == '*' 时
			// 只要前面的匹配，dp[0][j] 就匹配
			// 一旦 p[j-1] != '*'，后面的 dp[0][j] 就都为 false
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= sSize; i++ {
		for j := 1; j <= pSize; j++ {
			if p[j-1] != '*' {
				// 当 p[j-1] != '*' 时
				// 单个字符要匹配，并且之前的字符串也要匹配。
				dp[i][j] = dp[i-1][j-1] &&
					(p[j-1] == s[i-1] || p[j-1] == '?')
			} else {
				// 当 p[j-1] == '*' 时
				//   要么，dp[i-1][j] == true，意味着，
				//   当 s[:i] 与 p[:j+1] 匹配，且p[j] == '*' 的时候，
				//   s[:i] 后接任意字符串的 s[:i+1] 仍与 p[:j+1] 匹配。
				//   要么，dp[i][j-1] == true，意味着，
				//   当 s[:i+1] 与 p[:j] 匹配后
				//   在 p[:j] 后添加'*'，s[:i+1] 与 p[:j+1] 任然匹配
				//   因为， '*' 可以匹配空字符。
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[sSize][pSize]
}
