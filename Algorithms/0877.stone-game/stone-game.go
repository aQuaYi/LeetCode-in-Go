package problem0877

func stoneGame(p []int) bool {
	size := len(p)

	// dp[i][j] = k 表示, 面对 p[i:j+1] 个石头，先拿的人会比后拿的人多 k 个
	dp := [501][501]int{}
	for i := 0; i < size; i++ {
		dp[i][i] = p[i]
	}

	for d := 1; d < size; d++ {
		for i := 0; i < size-d; i++ {
			// p[i]-dp[i+1][i+d] 表示，如果拿了左边的会比后拿的人多的个数
			// p[i+d]-dp[i][i+d-1] 表示，如果拿了右边会比后拿的人多的个数
			// 注意 dp[i][i+d] 与 dp[i+1][i+d] (or dp[i][i+d-1]) 的先拿后拿的转换
			dp[i][i+d] = max(p[i]-dp[i+1][i+d], p[i+d]-dp[i][i+d-1])
		}
	}

	return dp[0][size-1] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/stone-game/discuss/154610/C++JavaPython-DP-or-Just-return-true
