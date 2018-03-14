package problem0685

// 本题存在 3 种情况需要处理
// 1. 存在一个节点，有两个父节点，但没有环状结构
// 2. 存在一个环状结构，且每个节点都只有一个父节点
// 3. 存在一个环状结构，且环上有且仅有一个节点有两个父节点

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)

	parent := make([]int, n+1)
	// first, last 用来保存指向同一个节点的两条 edge
	var first, last []int

	for k := range edges {
		p, c := edges[k][0], edges[k][1]
		if parent[c] == 0 {
			parent[c] = p
		} else {
			// c 是存在两个父节点的节点
			// 按照先后顺序保存这两个 edge 到 first 和 last
			first = []int{parent[c], c}
			last = edges[k]
			// edges[k] = nil 的用意是
			// 在 edges 中删除 last 边，然后查找环状结构
			// 如果 依然 存在环状结构的话，说明 first 边在环状结构上，需要删除的是 first 边
			// 如果  不  存在环状结构的话，不管 last 原先在不在环状结构上，都需要删除 last ，因为题目要求删除 occurs last 的边
			edges[k] = nil
			// 可以提前结束 for 循环了
			break
		}
	}

	// 以下是使用联通查找方法，查找环装结构
	root := parent

	for i := 0; i <= n; i++ {
		root[i] = i
	}

	rootOf := func(i int) int {
		for i != root[i] {
			i = root[i]
		}
		return i
	}

	for _, edge := range edges {
		if edge == nil {
			continue
		}

		p := edge[0]
		c := edge[1]
		r := rootOf(p)

		// r == c 说明，添加 edge 后，会形成一个封闭的环
		// edge 是让环封闭的边
		if r == c {
			if first == nil {
				// 不存在有两个父节点的节点
				// 返回让环封闭的 edge
				return edge
			}
			// 第 3 种情况出现了，且 first 边在环状结构上
			return first
		}

		// 按照 edge 出现的顺序，添加每个节点，目前的根节点
		root[c] = r
	}

	// 程序运行到这里，说明
	// 不存在环状结构
	// 那就删除 last 吧
	return last
}
