package problem0840

func numMagicSquaresInside(g [][]int) int {
	m, n := len(g), len(g[0])

	res := 0

	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			if !isAvailbleNum(i, j, g) {
				continue
			}

			if // 检查行
			g[i][j]+g[i][j+1]+g[i][j+2] == 15 &&
				g[i+1][j]+g[i+1][j+1]+g[i+1][j+2] == 15 &&
				g[i+2][j]+g[i+2][j+1]+g[i+2][j+2] == 15 &&
				// 检查列
				g[i][j]+g[i+1][j]+g[i+2][j] == 15 &&
				g[i][j+1]+g[i+1][j+1]+g[i+2][j+1] == 15 &&
				g[i][j+2]+g[i+1][j+2]+g[i+2][j+2] == 15 &&
				// 检查对角
				g[i][j]+g[i+1][j+1]+g[i+2][j+2] == 15 &&
				g[i][j+2]+g[i+1][j+1]+g[i+2][j] == 15 {
				res++
			}
		}
	}

	return res
}

func isAvailbleNum(x, y int, g [][]int) bool {
	tmp := [16]int{}

	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			tmp[g[i][j]]++
		}
	}

	for i := 1; i <= 9; i++ {
		if tmp[i] != 1 {
			return false
		}
	}
	return true
}
