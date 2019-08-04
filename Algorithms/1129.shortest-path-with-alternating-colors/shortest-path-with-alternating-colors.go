package problem1129

import (
	"math"
)

func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = math.MaxInt64
	}

	bfs := func(nodes []int, cur, next [][]int) {
		seenCur := make([]bool, n)
		seenNext := make([]bool, n)
		count := 0
		for len(nodes) > 0 {
			size := len(nodes)
			for k := 0; k < size; k++ {
				i := nodes[k]
				res[i] = min(res[i], count)
				for _, j := range cur[i] {
					if seenCur[j] {
						continue
					}
					seenCur[j] = true
					nodes = append(nodes, j)
				}
			}
			nodes = nodes[size:]
			count++
			// exchange colors
			cur, next = next, cur
			seenCur, seenNext = seenNext, seenCur
		}
	}

	red := trans(n, redEdges)
	blue := trans(n, blueEdges)

	bfs(make([]int, 1, n*2), red, blue)
	bfs(make([]int, 1, n*2), blue, red)

	for i := range res {
		if res[i] == math.MaxInt64 {
			res[i] = -1
		}
	}

	return res
}

// edge[i] contains all edges started form i
func trans(n int, edges [][]int) (edge [][]int) {
	edge = make([][]int, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		edge[i] = append(edge[i], j)
	}
	return edge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
