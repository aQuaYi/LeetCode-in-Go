package problem0847

func shortestPathLength(graph [][]int) int {
	size := len(graph)
	if size <= 3 {
		return size - 1
	}

	// 最短路径一定是从单连通的节点开始的。
	// 并且与同一个节点联通的单连通节点只需要添加一个
	// 这些候选的单连通节点放在 cands
	cands := make([]int, 0, size)
	isAdded := [13]bool{}
	for i, g := range graph {
		if len(g) == 1 && !isAdded[g[0]] {
			cands = append(cands, i)
			isAdded[g[0]] = true
		}
	}
	return 0
}
