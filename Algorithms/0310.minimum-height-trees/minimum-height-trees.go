package Problem0310

func findMinHeightTrees(n int, edges [][]int) []int {
	edgeMap := make([][]int, n)
	for i := range edgeMap {
		edgeMap[i] = make([]int, 0, n)
	}
	for _, e := range edges {
		edgeMap[e[0]] = append(edgeMap[e[0]], e[1])
		edgeMap[e[1]] = append(edgeMap[e[1]], e[0])
	}

	var res []int
	var rec []bool
	minHeight := 1<<31 - 1
	for i := 0; i < n; i++ {
		rec = make([]bool, n)
		nodes := make([]int, 1, n)
		nodes[0] = i
		rec[i] = true
		tempHeight := 1
		count := 1
		for len(nodes) > 0 {
			for _, node := range edgeMap[nodes[0]] {
				if !rec[node] {
					nodes = append(nodes, node)
					rec[node] = true
				}
			}
			nodes = nodes[1:]
			count--
			if count == 0 {
				tempHeight++
				if tempHeight > minHeight {
					break
				}
				count = len(nodes)
			}

		}

		if minHeight > tempHeight {
			minHeight = tempHeight
			res = []int{i}
		} else if minHeight == tempHeight {
			res = append(res, i)
		}

	}

	return res
}
