package Problem0174

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 1
	}

	n := len(dungeon[0])
	if n == 0 {
		return 1
	}

	// init[i][j] 到达 (i,j) 后还活着，需要的最小初始健康值
	init := make([][]int, m)
	// health[i][j] 到达 (i,j) 后的健康值
	health := make([][]int, m)
	for i := 0; i < m; i++ {
		init[i] = make([]int, n)
		health[i] = make([]int, n)
	}

	init[0][0] = 1
	health[0][0] = dungeon[0][0] + 1
	if dungeon[0][0] < 0 {
		init[0][0] = 1 - dungeon[0][0]
		health[0][0] = 1
	}

	for i := 1; i < m; i++ {
		health[i][0] = health[i-1][0] + dungeon[i][0]
		if health[i][0] < 1 {
			init[i][0] = init[i-1][0] - health[i][0] + 1
			health[i][0] = 1
		} else {
			init[i][0] = init[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		health[0][j] = health[0][j-1] + dungeon[0][j]
		if health[0][j] < 0 {
			init[0][j] = init[0][j-1] - health[0][j] + 1
			health[0][j] = 1
		} else {
			init[0][j] = init[0][j-1]
		}
	}
	iRight, hRight, iDown, hDown := 0, 0, 0, 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			iRight = init[i][j-1] - min(health[i][j-1]+dungeon[i][j]-1, 0)
			hRight = max(1, health[i][j-1]+dungeon[i][j])
			iDown = init[i-1][j] - min(health[i-1][j]+dungeon[i][j]-1, 0)
			hDown = max(1, health[i-1][j]+dungeon[i][j])
			switch {
			case iRight == iDown:
				init[i][j] = iRight
				health[i][j] = max(hRight, hDown)
			case iRight < iDown:
				init[i][j] = iRight
				health[i][j] = hRight
			default:
				init[i][j] = iDown
				health[i][j] = hDown
			}
		}
	}

	return init[m-1][n-1]
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
