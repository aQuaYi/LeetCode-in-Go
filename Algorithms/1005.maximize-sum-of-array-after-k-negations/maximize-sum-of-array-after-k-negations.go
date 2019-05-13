package problem1005

import "container/heap"

func largestSumAfterKNegations(A []int, K int) int {
	h := intHeap(A)
	heap.Init(&h)

	for K > 0 {
		h[0] = -h[0]
		heap.Fix(&h, 0)
		K--
	}

	sum := 0
	for i := 0; i < len(h); i++ {
		sum += h[i]
	}
	return sum
}

// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
