package problem0882

import (
	"sort"
)

const maxInt = 1<<62 - 1

func reachableNodes(edges [][]int, M int, N int) int {
	sort.Slice(edges, func(i int, j int) bool {
		if edges[i][0] == edges[j][0] {
			return edges[i][2] < edges[j][2]
		}
		return edges[i][0] < edges[j][0]
	})

	minDis := make([]int, N)
	for i := 1; i < N; i++ {
		minDis[i] = maxInt
	}
	minDis[0] = 0

	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		minDis[j] = min(minDis[j], minDis[i]+n+1)
		minDis[i] = min(minDis[i], minDis[j]+n+1)
	}

	res := totalEdgeNodes(edges) + N

	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		mdi, mdj := minDis[i], minDis[j]
		if mdi >= M && mdj >= M {
			res -= n
			continue
		}

		if mdi < M && mdj < M {
			res -= max(0, mdi+mdj+n-M*2)
			continue
		}

		md := min(mdi, mdj) // md < M

		res -= max(0, n-(M-md))
	}

	for _, d := range minDis {
		if d > M {
			res--
		}
	}

	return res
}

func totalEdgeNodes(edges [][]int) int {
	res := 0
	for _, e := range edges {
		res += e[2]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
