package Problem0684

func findRedundantConnection(edges [][]int) []int {
	N := len(edges)
	count := make([]int, N+1)
	var idx int
	max := 0
	for i, e := range edges {
		a, b := e[0], e[1]
		count[a]++
		count[b]++
		if max < count[a]+count[b] && count[a] > 1 && count[b] > 1 {
			max = count[a] + count[b]
			idx = i
		}
	}

	return edges[idx]
}
