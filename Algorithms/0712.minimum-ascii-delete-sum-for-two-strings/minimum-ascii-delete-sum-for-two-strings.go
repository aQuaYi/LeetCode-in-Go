package problem0712

func minimumDeleteSum(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)

	// dp[i][j] == miniumDeleteSum(s1[i:], s2[j:])
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := n1 - 1; 0 <= i; i-- {
		dp[i][n2] = dp[i+1][n2] + int(s1[i])
	}

	for j := n2 - 1; 0 <= j; j-- {
		dp[n1][j] = dp[n1][j+1] + int(s2[j])
	}

	for i := n1 - 1; 0 <= i; i-- {
		for j := n2 - 1; 0 <= j; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
