package problem1035

// ref: https://leetcode.com/problems/uncrossed-lines/discuss/282842/JavaC%2B%2BPython-DP-The-Longest-Common-Subsequence
func maxUncrossedLines(A []int, B []int) int {
	m, n := len(A), len(B)
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
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
