package problem0857

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	size := len(quality)

	workers := make([][2]float64, size)
	for i := 0; i < size; i++ {
		w, q := float64(wage[i]), float64(quality[i])
		workers[i] = [2]float64{w / q, q}
	}
	sort.Slice(workers, func(i int, j int) bool {
		return workers[i][0] < workers[j][0]
	})

	cost := math.MaxFloat64 // wageMax * qualityMax == 1e8
	qsum := 0.

	pq := make(PQ, 0, size)

	for _, w := range workers {
		qsum += w[1]
		heap.Push(&pq, w[1])

		if len(pq) > K {
			qsum -= heap.Pop(&pq).(float64)
		}

		if len(pq) == K {
			cost = min(cost, qsum*w[0])
		}
	}

	/**
	 * 为了满足题目的两个条件，可知，总的工资支出
	 * cost = (sum of k q) * (max of k w/q)
	 *
	 * 在 for 循环中， (max of k w/q) = w[0]
	 * 而 w[0] 是单调递增的，这一点很重要
	 *
	 * (sum of k q) = qsum
	 * 在 for 循环中， qsum 是单调递减的
	 *
	 * 由于 w[0] 和 qsum 的单调性
	 * 所以， 最小的 cost 注定在某个 qsum * w[0] 中
	 *
	 *
	 *
	 * 思考一下， 假设 workers[i][1] (i>K) 是最大的 quality
	 * 那么，qsum += workers[i][1] 后，立马就会 qsum -= workers[i][1]
	 * 也就是说，组成 qsum 的 k 个 worker 不包括 i
	 * 但是，还是会执行
	 * cost = min(cost, qsum*workers[i][0])
	 * 也就是说，(sum of k q) 与 (max of k w/q) 中的 k 不一致了，不会有问题吗？
	 *
	 *
	 *
	 * 答，不会有问题
	 * 在执行 min(cost, qsum*workers[i][0]) 的时候
	 * 可以肯定 cost <= qsum*workers[i-1][0]
	 * 因为，根据 w[0] 的单调性
	 * workers[i-1][0] <= workers[i][0]
	 * 也就是说 cost <= qsum*workers[i][0]
	 * 所以，
	 * qsum*workers[i][0] 虽然是一个错误的组合，但这个值不会带来错误的最小值
	 *
	 */

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
