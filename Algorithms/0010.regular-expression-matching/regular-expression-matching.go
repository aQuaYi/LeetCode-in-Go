package problem0010

// 程序中存在以下假设
// "*" 不会出现在p的首位
// "**" 不会出现，但会出现 ".*."" , ".*.." , ".*.*"

func isMatch(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for i := 1; i < len(p); i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
