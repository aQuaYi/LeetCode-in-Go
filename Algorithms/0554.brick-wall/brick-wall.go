package Problem0554

func leastBricks(wall [][]int) int {
	m := len(wall)

	rec := make(map[int]int, m)
	max := 0

	for i := 0; i < m; i++ {
		sum := wall[i][0]
		for j := 1; j < len(wall[i]); j++ {
			rec[sum]++
			if max < rec[sum] {
				max = rec[sum]
			}
			sum += wall[i][j]
		}
	}

	return m - max
}
