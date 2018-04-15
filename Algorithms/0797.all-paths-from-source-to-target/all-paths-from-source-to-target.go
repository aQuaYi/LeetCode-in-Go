package problem0797

func allPathsSourceTarget(graph [][]int) [][]int {
	path := make([]int, len(graph))
	res := make([][]int, 0, len(graph))

	dfs(0, len(graph)-1, 1, path, graph, &res)

	return res
}

func dfs(id, dst, pathLen int, path []int, graph [][]int, res *[][]int) {
	if id == dst {
		temp := make([]int, pathLen)
		copy(temp, path)
		*res = append(*res, temp)
	}

	for _, node := range graph[id] {
		path[pathLen] = node
		dfs(node, dst, pathLen+1, path, graph, res)
	}
}
