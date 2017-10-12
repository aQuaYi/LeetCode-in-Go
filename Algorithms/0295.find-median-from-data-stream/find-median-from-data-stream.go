package Problem0295

import (
	"container/heap"
)

// MedianFinder 用于寻找 Median
type MedianFinder struct {
	left  *maxHeap
	right *minHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	left := new(maxHeap)
	heap.Init(left)
	right := new(minHeap)
	heap.Init(right)

	return MedianFinder{
		left:  left,
		right: right,
	}
}

// AddNum 给 MedianFinder 添加数
func (mf *MedianFinder) AddNum(n int) {
	if mf.left.Len() == mf.right.Len() {
		heap.Push(mf.left, n)
	} else {
		heap.Push(mf.right, n)
	}

	if mf.right.Len() > 0 && mf.left.intHeap[0] > mf.right.intHeap[0] {
		l := heap.Pop(mf.left)
		r := heap.Pop(mf.right)
		heap.Push(mf.left, r)
		heap.Push(mf.right, l)
	}
}

// FindMedian 给出 Median
func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() == mf.right.Len() {
		return float64(mf.left.intHeap[0]+mf.right.intHeap[0]) / 2
	}
	return float64(mf.left.intHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type maxHeap struct {
	intHeap
}

func (h maxHeap) Less(i, j int) bool {
	return h.intHeap[i] > h.intHeap[j]
}

type minHeap struct {
	intHeap
}

func (h minHeap) Less(i, j int) bool {
	return h.intHeap[i] < h.intHeap[j]
}

// myHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
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
	*h = (*h)[0 : len(*h)-1]
	return res
}
