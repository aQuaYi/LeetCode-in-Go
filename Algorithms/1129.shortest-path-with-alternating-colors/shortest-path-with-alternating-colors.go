package problem1129

import (
	"math"
)

func shortestAlternatingPaths(n int, reds [][]int, blues [][]int) []int {
	hasRedEdge := convert(n, reds)
	hasBlueEdge := convert(n, blues)

	res := make([]int, n)
	for i := range res {
		res[i] = math.MaxInt64
	}

	bfs := func(g []int, hasCur, hasNext [][]bool) {
		seenCur := make([]bool, n)
		seenNext := make([]bool, n)
		count := 0
		for len(g) > 0 {
			size := len(g)
			for k := 0; k < size; k++ {
				i := g[k]
				res[i] = min(res[i], count)
				for j := 0; j < n; j++ {
					if seenCur[j] || !hasCur[i][j] {
						continue
					}
					seenCur[j] = true
					g = append(g, j)
				}
			}
			g = g[size:]
			count++
			// exchange colors
			hasCur, hasNext = hasNext, hasCur
			seenCur, seenNext = seenNext, seenCur
		}
	}

	bfs(make([]int, 1, n*2), hasRedEdge, hasBlueEdge)
	bfs(make([]int, 1, n*2), hasBlueEdge, hasRedEdge)

	for i := range res {
		if res[i] == math.MaxInt64 {
			res[i] = -1
		}
	}

	return res
}

// hasEdge[i][j]==true means has a directed edge form i to j
func convert(n int, edges [][]int) (hasEdge [][]bool) {
	hasEdge = make([][]bool, n)
	for i := range hasEdge {
		hasEdge[i] = make([]bool, n)
	}
	for _, e := range edges {
		i, j := e[0], e[1]
		hasEdge[i][j] = true
	}
	return hasEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
