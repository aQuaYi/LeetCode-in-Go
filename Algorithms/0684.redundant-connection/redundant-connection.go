package Problem0684

func findRedundantConnection(edges [][]int) []int {
	N := len(edges)
	// connected
	conn := make([][]int, N+1)
	triple := 0
	for _, e := range edges {
		conn[e[0]] = append(conn[e[0]], e[1])
		conn[e[1]] = append(conn[e[1]], e[0])
		if len(conn[e[0]]) == 3 {
			triple = e[0]
		} else if len(conn[e[1]]) == 3 {
			triple = e[1]
		}
	}

	if triple == 0 {
		return edges[N-1]
	}

	var node int
	for i := 2; 0 <= i; i-- {
		node = conn[triple][i]
		if len(conn[node]) == 2 {
			break
		}
	}
	return []int{triple, node}
}
