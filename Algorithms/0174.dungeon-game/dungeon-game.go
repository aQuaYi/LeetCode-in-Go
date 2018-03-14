package problem0174

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 1
	}

	n := len(dungeon[0])
	if n == 0 {
		return 1
	}

	intMax := 1<<63 - 1
	// dp[i][j] 到达 (i,j) 点**前**的健康值
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			// (i,j) 点的健康值达到上限，可以保证到达 (m-1, n-1) 点
			dp[i][j] = intMax
		}
	}

	// health 是到达 (i,j) 点**前**的健康值
	health := 0
	// 把 dp[m][n-1] = 1 更换成 dp[m-1][n] = 1，程序一样可以通过
	// 可以这样认为
	// 还存在一个 dungeon[m][n-1] == 0 的房间
	// 骑士救出公主后，向下移动一格，达到这个 (m, n-1) 房间
	// 骑士在这个房间的健康值 只需要为 1 就可以了。
	dp[m][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// health + dungeon[i][j] == min(dp[i+1][j], dp[i][j+1]) 的意思是
			// 骑士达到 (i,j) 点**前**的健康值，
			// 要能够保证他至少移动到 右边 或者 左边 房间中的一个
			health = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			// 由于 dungeon[i][j] 也可能是正数
			// 可能会导致 health < 1 。此时，骑士无法到达 (i,j) 房间，这不符合题意
			// 因此，当 health < 1 时,
			// dp[i][j] = 1
			dp[i][j] = max(health, 1)
		}
	}

	// dp[0][0] 是骑士进入 (0,0) 房间前的健康值
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
