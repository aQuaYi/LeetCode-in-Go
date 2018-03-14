package problem0062

func uniquePaths(m int, n int) int {
	// dp[i][j] 代表了，到达 (i,j) 格子的不同路径数目
	dp := [][]int{}

	// 创建棋盘
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}

	for i := 0; i < m; i++ {
		// 到达第 0 列的格子，只有一条路径
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// 到达第 0 行的格子，只有一条路径
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达 (i,j) 格子的路径数目，等于
			// 到达 上方格子 和 左边格子 路径数之和
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
