package problem0785

func isBipartite(graph [][]int) bool {
	// set[i] == 'a' 表示 node i 在 a 集合中
	// set[i] == 'b' 表示 node i 在 b 集合中
	// set[i] == 0 表示 node i 还没有被标记集合
	set := make([]byte, len(graph))

	for i := 0; i < len(graph); i++ {
		// len(graph[i])>0 说明，node i 与其他 node 相互联通
		// set[i]== 0 说明，所有与 node i 相互联通的点，都没有被检查过
		if len(graph[i]) > 0 && set[i] == 0 {
			if !isOK(i, set, graph) {
				return false
			}
		}
	}

	return true
}

func isOK(i int, set []byte, graph [][]int) bool {
	// 把 node i 放入 a 集合
	set[i] = 'a'
	// nodes 收集了同在一个 set 中的所有点
	nodes := []int{i}
	// 与 nodes 所在集合相反的集合名称
	theOtherSet := byte('b')
	otherNodes := make([]int, 0, 100)

	for {
		if len(nodes) == 0 {
			// 当前集合中的 node 检查完了
			// 就去检查另一边的集合中的 node
			nodes = otherNodes
			if len(nodes) == 0 {
				// 另一边集合中，没有 node
				// 说明，全部检查完了
				break
			}
			// 为 other nodes 提供存在容器
			otherNodes = make([]int, 0, 100)
			// 变化集合名称
			theOtherSet = change(theOtherSet)
		}

		// 逐个检查与 nodes[0] 相连的点
		for _, nd := range graph[nodes[0]] {
			// 如果 nd 已经被分好类了
			if set[nd] == 'a' || set[nd] == 'b' {
				// 检查 nd 的分类会不会出现矛盾
				if set[nd] != theOtherSet {
					// 出现矛盾说明，不能满足题意
					return false
				}
			} else { // 还没有分好类
				// 就为 nd 标记分类
				set[nd] = theOtherSet
				// 并放入 otherNodes 当中
				otherNodes = append(otherNodes, nd)
			}
		}

		// 为了 for 循环可以结束
		// 从 nodes 中排除 nodes[0]
		nodes = nodes[1:]
	}

	// 全部检查完成，没有出错
	// 说明满足题意要求
	return true
}

func change(s byte) byte {
	if s == 'a' {
		return 'b'
	}
	return 'a'
}
