package problem0973

import "container/heap"

func kClosest(points [][]int, K int) [][]int {
	res := make([][]int, 0, K)
	h := intsHeap(points)
	heap.Init(&h)
	for K > 0 {
		res = append(res, heap.Pop(&h).([]int))
		K--
	}
	return res
}

// intsHeap 实现了 heap 的接口
type intsHeap [][]int

func (h intsHeap) Len() int {
	return len(h)
}

func (h intsHeap) Less(i, j int) bool {
	return length(h[i]) < length(h[j])
}

func (h intsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intsHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.([]int))
}

func (h *intsHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func length(a []int) int {
	return a[0]*a[0] + a[1]*a[1]
}
