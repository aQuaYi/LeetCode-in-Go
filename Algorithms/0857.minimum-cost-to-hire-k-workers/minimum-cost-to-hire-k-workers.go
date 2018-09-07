package problem0857

import (
	"container/heap"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	size := len(quality)

	workers := make([][2]float64, size)
	for i := 0; i < size; i++ {
		w, q := float64(wage[i]), float64(quality[i])
		ratio := w / q
		workers[i] = [2]float64{ratio, q}
	}
	sort.Slice(workers, func(i int, j int) bool {
		return workers[i][0] < workers[j][0]
	})

	pq := make(PQ, 0, size)
	qSum := 0.
	for i := 0; i < K; i++ {
		q := workers[i][1]
		pq = append(pq, q)
		qSum += q
	}

	maxRatio := workers[K-1][0]

	cost := qSum * maxRatio

	heap.Init(&pq)

	for i := K; i < size; i++ {
		maxRatio, q := workers[i][0], workers[i][1]
		if q >= pq[0] {
			/* q >= pq[0] 时，qSum 不变，maxRatio 变大，qSum*maxRatio 不会是新低 */
			continue
		}
		heap.Push(&pq, q)
		qSum += q
		qSum -= heap.Pop(&pq).(float64)
		/* qSum 总是 workers[:i+1] 中 K 个最小的 q 之和 */
		cost = min(cost, qSum*maxRatio)
	}

	return cost
}

// PQ implements heap.Interface
type PQ []float64

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 x
func (pq *PQ) Push(x interface{}) {
	temp := x.(float64)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 x
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
