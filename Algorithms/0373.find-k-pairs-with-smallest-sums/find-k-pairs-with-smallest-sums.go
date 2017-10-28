package Problem0373

import "container/heap"

type pair struct {
	i     int
	j     int
	sum   int
	index int
}

type priorityQueue []*pair

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	p := x.(*pair)
	p.index = n
	*pq = append(*pq, p)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	p := old[n-1]
	p.index = -1
	*pq = old[0 : n-1]
	return p
}

func kSmallestPairs(a, b []int, k int) [][]int {
	var res [][]int
	if len(a) == 0 || len(b) == 0 {
		return res
	}

	pqLen := min(k, len(a))

	// 先把每个 a[i] 与 b[0] 结合，放入 pq
	pq := make(priorityQueue, pqLen)
	for l := 0; l < k && l < len(a); l++ {
		pq[l] = &pair{i: l, j: 0, sum: a[l] + b[0], index: l}
	}
	// 初始化 pq
	heap.Init(&pq)

	for m := k; m > 0 && len(pq) > 0; m-- {
		current := heap.Pop(&pq).(*pair)
		res = append(res, []int{a[current.i], b[current.j]})
		if current.j+1 < len(b) {
			heap.Push(&pq, &pair{i: current.i, j: current.j + 1, sum: a[current.i] + b[current.j+1]})
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
