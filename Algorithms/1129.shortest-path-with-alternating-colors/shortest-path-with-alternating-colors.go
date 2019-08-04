package problem1129

func shortestAlternatingPaths(n int, reds [][]int, blues [][]int) []int {
	isRed := convert(n, reds)
	isBlue := convert(n, blues)

	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	bfs := func(g []int, isCur, isNext [][]bool) {
		hasSeen := make([]bool, n)
		count := 0
		for len(g) > 0 {
			size := len(g)
			for k := 0; k < size; k++ {
				i := g[k]
				res[i] = min(res[i], count)
				hasSeen[i] = true
				for j := 0; j < n; j++ {
					if hasSeen[j] || !isCur[i][j] {
						continue
					}
					g = append(g, j)
				}
			}
			g = g[size:]
			isCur, isNext = isNext, isCur
			count++
		}
	}

	bfs([]int{0}, isRed, isBlue)
	bfs([]int{0}, isBlue, isRed)

	for i := range res {
		if res[i] == n {
			res[i] = -1
		}
	}

	return res
}

func convert(n int, edges [][]int) [][]bool {
	res := make([][]bool, n)
	for i := range res {
		res[i] = make([]bool, n)
	}
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
