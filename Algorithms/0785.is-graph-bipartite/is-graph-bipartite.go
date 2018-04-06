package problem0785

func isBipartite(graph [][]int) bool {
	n := len(graph)
	// painted[i] == 1  表示 node i 在 白色 集合中
	// painted[i] == -1 表示 node i 在 黑色 集合中
	// painted[i] == 0  表示 node i 还没有被标记集合
	painted := make([]int, n)

	for i := 0; i < n; i++ {
		// painted[i]== 0 说明，所有与 node i 相互联通的点，都没有被检查过
		if painted[i] == 0 && !isOK(i, 1, painted, graph) {
			return false
		}
	}

	return true
}

// isOK 会把 node 归类为 color 集合
// 在符合题意时，返回 true
func isOK(node, color int, painted []int, graph [][]int) bool {
	// 如果 node 已经被检查过了
	// 可以直接核对颜色
	if painted[node] != 0 {
		return painted[node] == color
	}

	// 按照要求，对 node 进行归类
	painted[node] = color

	// 对与 node 连接的点，也进行同样的检查
	for _, i := range graph[node] {
		// 注意，转变 color
		if !isOK(i, -color, painted, graph) {
			return false
		}
	}

	return true
}
