package problem0474

import (
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 为 i 个 '0' 和 j 个 '1' 所能形成的最多的 s 的个数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for i := m; i-zeros >= 0; i-- {
			for j := n; j-ones >= 0; j-- {
				// 更新 dp[i][j]
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
