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

	min := 1<<63 - 1

	var dfs func(int, int, int, int)
	// health 是达到 (i,j) 点前的健康值
	// init 是到达 (0,0) 点前的初始值
	dfs = func(i, j, health, init int) {
		if i < m && j < n {
			//,到达 (i,j) 后，更新 health 值
			health += dungeon[i][j]
			if health < 1 {
				// 提升 init 值
				init = init - health + 1
				// 确保到达 (i,j) 后，health == 1
				health = 1
			}

			if i == m-1 && j == n-1 && min > init {
				// 到达右下角，且出现新低
				min = init
			}
		}

		if i+1 < m {
			// 还能往下走，就往下走一步
			dfs(i+1, j, health, init)
		}

		if j+1 < n {
			// 还能往右走，就往右走一步
			dfs(i, j+1, health, init)
		}
	}

	dfs(0, 0, 1, 1)

	return min
}
