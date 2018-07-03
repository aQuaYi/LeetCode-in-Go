package problem0847

func shortestPathLength(graph [][]int) int {
	size := len(graph)
	maskSize := 1 << uint(size)
	maxInt := 1<<63 - 1

	queue := make([]state, 0, size*size)
	dp := make([][]int, size)
	// dp[3][11(00...001011)] = 2 的含义是
	// 从 3 出发，包含了 0,1,3 点的路径距离是 2
	// 通过把二进制数 0,1,3 位上的值设置为 1，其余位为 0 ，来表示包含了 0，1，3 点的路径
	// 11 是其二进制 1011 的 mask

	for i := range dp {
		dp[i] = make([]int, maskSize)
		for j := range dp[i] {
			// 先把每个路径都设置成最大值
			dp[i][j] = maxInt
		}
		mask := 1 << uint(i)
		// 从自身出发到达自身的距离为 0
		dp[i][mask] = 0
		// 从每个节点出发都可以作为一个状态
		queue = append(queue, state{source: i, mask: mask})
	}

	// 搜索
	for len(queue) > 0 {
		// 从 queue 中获取一个状态
		s := queue[0]
		queue = queue[1:]

		// 检查从 s 状态的 source 能够到达的所有节点
		// 是否有新的更优状态产生
		// 有的话，就放入 queue 中

		for _, next := range graph[s.source] {
			// nextMask 很巧妙
			// nextMask 表示，s 状态所包含的节点，并上 next 节点
			nextMask := s.mask | 1<<uint(next)
			if dp[next][nextMask] > dp[s.source][s.mask]+1 {
				dp[next][nextMask] = dp[s.source][s.mask] + 1
				// 注意，只有更优的状态才能放入 queue 中，避免了死循环
				queue = append(queue, state{source: next, mask: nextMask})
			}
		}
	}

	res := maxInt
	traversalMask := maskSize - 1
	for i := 0; i < size; i++ {
		res = min(res, dp[i][traversalMask])
	}

	return res
}

type state struct {
	source, mask int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
