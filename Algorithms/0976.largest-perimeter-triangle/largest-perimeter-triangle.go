package problem0976

import "container/heap"

func largestPerimeter(A []int) int {
	size := len(A)
	h := intHeap(A)

	heap.Init(&h)

	a := heap.Pop(&h).(int)
	b := heap.Pop(&h).(int)
	for i := size - 3; i >= 0; i-- {
		c := heap.Pop(&h).(int)
		if a < b+c {
			return a + b + c
		}
		a, b = b, c
	}
	return 0
}

// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j] // NOTICE: Max is at the top
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
