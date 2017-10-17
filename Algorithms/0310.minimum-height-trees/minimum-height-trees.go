package Problem0310

func findMinHeightTrees(n int, edges [][]int) []int {
	eMap := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		eMap[i] = make([]int, 0, n)
	}

	for _, e := range edges {
		eMap[e[0]] = append(eMap[e[0]], e[1])
		eMap[e[1]] = append(eMap[e[1]], e[0])
	}

	var ok bool
	var nodes []int
	var i, nd int

	for len(eMap) > 2 {
		for i = range eMap {
			if len(eMap[i]) == 1 {
				delete(eMap, i)
			}
		}

		for i, nodes = range eMap {
			newNodes := make([]int, 0, len(nodes))
			for _, nd = range nodes {
				if _, ok = eMap[nd]; ok {
					newNodes = append(newNodes, nd)
				}
				eMap[i] = newNodes
			}
		}
	}

	var res = make([]int, 0, len(eMap))
	for k := range eMap {
		res = append(res, k)
	}

	return res
}
