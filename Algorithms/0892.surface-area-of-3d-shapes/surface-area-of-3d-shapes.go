package problem0892

func surfaceArea(grid [][]int) int {
	N := len(grid)
	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}
			res += v*6 - (v-1)*2 /** 单独一个立柱的外表面面积 */
			if j+1 < N {
				res -= min(v, grid[i][j+1]) * 2 /** 与右侧立柱贴合的面积 */
			}
			if i+1 < N {
				res -= min(v, grid[i+1][j]) * 2 /** 与下侧立柱贴合的面积 */
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
