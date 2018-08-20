package problem0882

import "container/heap"

func reachableNodes(edges [][]int, M int, N int) int {
	nodes := make(map[int]int, len(edges))
	nextTo := make([][]int, N)
	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		nodes[encode(i, j)] = n
		nextTo[i] = append(nextTo[i], j)
		nextTo[j] = append(nextTo[j], i)
	}

	// pq[x] = []int{r, i} 表示，
	// 到达 i 节点时，还可以走 r 步
	pq := make(PQ, 1, 1000)
	pq[0] = []int{M, 0}

	// pq 很有可能存在多个 {r1, i}, {r2, i}, {r3, i}
	// 利用 pq 的特性，优先处理 r 值最大的情况。
	// 并在处理时，
	// 在 seen 中标记已经处理过 i
	// 在 maxMovesRemaining 中记录 r 的最大值
	seen := [3001]bool{}
	// maxMovesRemaining[3] = 10 表示，
	// 通过最短的路径到达 3 节点时，还可以走 10 步
	maxMovesRemaining := [3001]int{}

	res := 0
	for len(pq) > 0 {
		r := pq[0][0]
		i := pq[0][1]
		heap.Pop(&pq)
		if seen[i] {
			continue
		}
		seen[i] = true
		maxMovesRemaining[i] = r
		res++ // 收获 edge 端点 i

		for _, j := range nextTo[i] {
			if seen[j] {
				continue
			}
			n := nodes[encode(i, j)]
			jRemaining := r - n - 1
			if jRemaining >= 0 {
				// 如果可以在 M 步内到达 j
				heap.Push(&pq, []int{jRemaining, j})
			}
		}
	}

	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		// maxMovesRemaining[i] = ri 表示达到 i 点后，最多还可以走 ri 步
		// maxMovesRemaining[j] = rj 表示达到 j 点后，最多还可以走 rj 步
		// 如果 ri + rj >= n, 则 edge(i,j) 中间的 n 个点都可以被走到
		// 否则 edge(i,j) 中只有 ri+rj 个点被走到
		res += min(maxMovesRemaining[i]+maxMovesRemaining[j], n)
	}

	return res
}

func encode(i, j int) int {
	if i > j {
		i, j = j, i
	}
	return i<<16 | j
}

// PQ implements heap.Interface and holds entries.
type PQ [][]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.([]int)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
