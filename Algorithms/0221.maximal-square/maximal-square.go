package problem0221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	maxEdge := 0
	// dp[i][j] == 以 (i,j) 点为右下角点的符合题意的最大正方形的边长
	// TODO: 由于 dp[i][j] 只与上，左上，左的数据有关，可以把 dp 压缩成一维的
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				maxEdge = max(maxEdge, dp[i][j])
			}
		}
	}

	return maxEdge * maxEdge
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
