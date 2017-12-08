package Problem0685

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parentCount := make([]int, n+1)
	childrenOf := make([][]int, n+1)

	for i := 0; i < n; i++ {
		u, v := edges[i][0], edges[i][1]

		parentCount[v]++
		if parentCount[v] > 1 {
			return edges[i]
		}

		childrenOf[u] = append(childrenOf[u], v)
	}

	isInCircle := func(edge []int) bool {
		p, c := edge[0], edge[1]
		queue := make([]int, len(childrenOf[c]), n)
		copy(queue, childrenOf[c])

		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]

			if x == p {
				return true
			}

			queue = append(queue, childrenOf[x]...)
		}

		return false
	}

	i := n - 1
	for ; 0 <= i; i-- {
		if isInCircle(edges[i]) {
			break
		}
	}

	return edges[i]
}
