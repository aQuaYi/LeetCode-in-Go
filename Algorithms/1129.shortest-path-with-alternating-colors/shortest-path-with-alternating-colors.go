package problem1129

func shortestAlternatingPaths(n int, reds [][]int, blues [][]int) []int {
	R := convert(n, reds)
	B := convert(n, blues)

	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	bfs := func(g []int, isRed, isBlue [][]bool) {
		hasSeen := make([]bool, n)
		count := 0
		for len(g) > 0 {
			size := len(g)
			for k := 0; k < size; k++ {
				i := g[k]
				res[i] = min(res[i], count)
				hasSeen[i] = true
				for j := 0; j < n; j++ {
					if hasSeen[j] || !isRed[i][j] {
						continue
					}
					g = append(g, j)
				}
			}
			g = g[size:]
			isRed, isBlue = isBlue, isRed
		}
	}

	bfs([]int{0}, R, B)
	bfs([]int{0}, B, R)

	return res
}

func convert(n int, edges [][]int) [][]bool {
	res := make([][]bool, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		res[i][j] = true
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
