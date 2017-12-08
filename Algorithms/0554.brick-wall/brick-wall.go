package Problem0554

func leastBricks(wall [][]int) int {
	m := len(wall)

	rec := make(map[int]int, m)

	for i := 0; i < m; i++ {
		sum := wall[i][0]
		for j := 1; j < len(wall[i]); j++ {
			rec[sum]++
			sum += wall[i][j]
		}
	}

	max := 0
	for _, bricks := range rec {
		if max < bricks {
			max = bricks
		}
	}

	return m - max
}
