package Problem0685

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parentCount := make([]int, n+1)
	childrenOf := make([][]int, n+1)
	multiParent := 0

	for i := 0; i < n; i++ {
		u, v := edges[i][0], edges[i][1]

		parentCount[v]++
		if parentCount[v] > 1 {
			multiParent = v
		}

		childrenOf[u] = append(childrenOf[u], v)
	}

	isInCircle := func(edge []int) bool {
		p, c := edge[0], edge[1]
		queue := make([]int, n)
		isAdded := make([]bool, n+1)

		for _, cc := range childrenOf[c] {
			queue = append(queue, cc)
			isAdded[cc] = true
		}

		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]

			if x == p {
				return true
			}

			for _, cc := range childrenOf[x] {
				if !isAdded[cc] {
					queue = append(queue, cc)
					isAdded[cc] = true
				}
			}

		}

		return false
	}

	hasCircle := false
	for i := 0; i < n; i++ {
		if isInCircle(edges[i]) {
			hasCircle = true
			break
		}
	}

	if multiParent > 0 && hasCircle {
		for i := n - 1; 0 <= i; i-- {
			if multiParent != edges[i][1] {
				continue
			}
			if isInCircle(edges[i]) {
				return edges[i]
			}
		}
	}

	if multiParent > 0 {
		for i := n - 1; 0 <= i; i-- {
			if edges[i][1] == multiParent {
				return edges[i]
			}
		}
	}

	i := n - 1
	for ; 0 <= i; i-- {
		if isInCircle(edges[i]) {
			break
		}
	}
	return edges[i]
}
