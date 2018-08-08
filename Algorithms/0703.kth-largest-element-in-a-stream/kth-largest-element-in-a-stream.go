package problem0703

import (
	"container/heap"
)

// KthLargest object will be instantiated and called as such:
// obj := Constructor(k, nums);
// param_1 := obj.Add(val);
type KthLargest struct {
	k    int
	heap intHeap
}

// Constructor 创建 KthLargest
func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

// Add 负责添加元素
func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	if len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

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
