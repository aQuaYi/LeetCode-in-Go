package problem0882

import "container/heap"

func reachableNodes(edges [][]int, M int, N int) int {
	nodes := make(map[int]int, len(edges))
	connects := make([][]int, N)
	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		nodes[encode(i, j)] = n
		connects[i] = append(connects[i], j)
		connects[j] = append(connects[j], i)
	}

	pq := make(PQ, 1, 1000)
	pq[0] = []int{M, 0}
	seen := make(map[int]int, 1000)

	for len(pq) > 0 {
		moves := pq[0][0]
		i := pq[0][1]
		heap.Pop(&pq)
		if _, ok := seen[i]; !ok {
			seen[i] = moves
			for _, j := range connects[i] {
				if _, ok := seen[j]; ok {
					continue
				}
				n := nodes[encode(i, j)]
				jMoves := moves - n - 1
				if jMoves >= 0 {
					heap.Push(&pq, []int{jMoves, j})
				}
			}
		}
	}

	res := len(seen)
	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		res += min(seen[i]+seen[j], n)
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
	// if pq[i][0] < pq[j][0]{
	// return pq[i][1] < pq[j][1]
	// }
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
