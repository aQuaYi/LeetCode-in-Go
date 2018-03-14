package problem0373

import "container/heap"

type pair struct {
	i   int
	j   int
	sum int
}

type priorityQueue []*pair

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	p := x.(*pair)
	*pq = append(*pq, p)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	p := old[n-1]
	*pq = old[0 : n-1]
	return p
}

func kSmallestPairs(a, b []int, k int) [][]int {
	var res [][]int
	if len(a) == 0 || len(b) == 0 {
		return res
	}

	// 可以想象一下，存在一个 len(a) * len(b) 的矩阵 mat
	// mat[i][j]  == a[i]+b[j]
	// res 就是 mat 中前 k 项值的坐标对
	// 由题意可知，a，b 递增，那么 mat 中的每一行和每一列也是递增的。

	pqLen := min(k, len(a))
	// 先把 mat[:][0] 的值放入 pq
	pq := make(priorityQueue, pqLen)
	for l := 0; l < k && l < len(a); l++ {
		pq[l] = &pair{i: l, j: 0, sum: a[l] + b[0]}
	}
	// 初始化 pq
	heap.Init(&pq)

	var min *pair

	for ; k > 0 && len(pq) > 0; k-- {
		// 获取 heap 中 sum 值最小的 pair
		min = heap.Pop(&pq).(*pair)
		// 加入到 res
		res = append(res, []int{a[min.i], b[min.j]})
		// mat[i][j] 被 pop 出去了，就把 mat[i][j+1] push 到 pq
		// 保证 mat 中每一行都有一个最小的在 pq 中，
		// 就可以保证 pq 中的 min 就是下一个 sum值最小的元素
		if min.j+1 < len(b) {
			heap.Push(&pq, &pair{i: min.i, j: min.j + 1, sum: a[min.i] + b[min.j+1]})
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
