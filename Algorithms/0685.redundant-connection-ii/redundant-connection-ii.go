package Problem0685

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)

	parent := make([]int, n+1)
	canA := make([]int, 0)
	canB := make([]int, 0)

	for k, v := range edges {
		if parent[v[1]] == 0 {
			parent[v[1]] = v[0]
		} else {
			canA = []int{parent[v[1]], v[1]}
			canB = v
			tmp := []int{v[0], 0}
			edges[k] = tmp
		}
	}

	root := func(i int) int {
		for i != parent[i] {
			i = parent[i]
		}
		return i
	}

	for i := 0; i <= n; i++ {
		parent[i] = i
	}

	for _, v := range edges {
		if v[1] == 0 {
			continue
		}

		p := v[0]
		c := v[1]
		a := root(p)

		if a == c {
			if len(canA) > 0 {
				return canA
			}
			return v
		}
		parent[c] = a
	}
	return canB
}
