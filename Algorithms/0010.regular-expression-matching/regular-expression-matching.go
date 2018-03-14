package problem0010

// 程序中存在以下假设
// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true

	for i := 2; i < len(dp); i += 2 {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	for i := 1; i < len(dp); i++ {
		if i < len(p) && p[i] == '*' {
			continue
		}

		for j := 1; j < len(dp[0]); j++ {

			if p[i-1] == '*' {
				if p[i-2] == '.' {
					dp[i][j] = dp[i-2][j-1] || dp[i][j-1] || dp[i-2][j]
				} else {
					dp[i][j] = (dp[i-2][j]) || (p[i-2] == s[j-1] && (dp[i-2][j-1] || dp[i][j-1]))
				}
			} else if p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
			}
		}
	}

	return dp[len(p)][len(s)]
}
