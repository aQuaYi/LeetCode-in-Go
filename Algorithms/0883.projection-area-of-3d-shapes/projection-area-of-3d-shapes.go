package problem0883

func projectionArea(grids [][]int) int {
	xy := make(map[int]bool, 2500)
	yz := make(map[int]bool, 2500)
	xz := make(map[int]bool, 2500)

	for i, grid := range grids {
		for j, cubes := range grid {
			for k := 0; k < cubes; k++ {
				xy[encode(i, j)] = true
				yz[encode(j, k)] = true
				xz[encode(i, k)] = true
			}
		}
	}

	return len(xy) + len(yz) + len(xz)
}

func encode(a, b int) int {
	return a<<8 | b
}
