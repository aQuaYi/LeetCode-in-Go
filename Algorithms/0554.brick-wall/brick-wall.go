package problem0554

func leastBricks(wall [][]int) int {
	m := len(wall)

	// 统计坐标相同的 edge 出现的次数
	count := make(map[int]int, m)

	for i := 0; i < m; i++ {
		sum := wall[i][0]
		for j := 1; j < len(wall[i]); j++ {
			count[sum]++
			sum += wall[i][j]
		}
	}

	max := 0
	for _, edges := range count {
		if max < edges {
			max = edges
		}
	}

	return m - max
}
