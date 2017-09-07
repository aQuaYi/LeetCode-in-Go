package Problem0087

func isScramble(a, b string) bool {
	if a == b {
		return true
	}
	// 题目规定了， s1 与 s2 等长
	n := len(a)

	dp := make([][][]bool, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = make([][]bool, n)
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = make([]bool, n+1)

			maxK := n - max(i, j)
			for k := 1; k <= maxK; k++ {
				if a[i:i+k] == b[j:j+k] {
					dp[i][j][k] = true
				} else {
					for d := 1; d < k; d++ {
						if (dp[i][j][d] && dp[i+d][j+d][k-d]) ||
							(dp[i][j+k-d][d] && dp[i+d][j][k-d]) {
							dp[i][j][k] = true
							break
						}
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
