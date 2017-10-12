package Problem0295

import (
	"container/heap"
)

// MedianFinder 用于寻找 Median
type MedianFinder struct {
	isEven bool
	left   *leftHeap
	right  *rightHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	isEven := true
	left := new(leftHeap)
	heap.Init(left)
	right := new(rightHeap)
	heap.Init(right)

	return MedianFinder{
		isEven: isEven,
		left:   left,
		right:  right,
	}
}

// AddNum 给 MedianFinder 添加数
func (mf *MedianFinder) AddNum(n int) {

	if mf.isEven {
		if len(mf.left.myHeap) == 0 {
			heap.Push(mf.left, n)
		} else if n <= mf.right.myHeap[0] {
			heap.Push(mf.left, n)
		} else {
			heap.Push(mf.left, mf.right.myHeap[0])
			mf.right.myHeap[0] = n
			heap.Fix(mf.right, 0)
		}
	} else {
		if mf.left.myHeap[0] <= n {
			heap.Push(mf.right, n)
		} else {
			heap.Push(mf.right, mf.left.myHeap[0])
			mf.left.myHeap[0] = n
			heap.Fix(mf.left, 0)
		}
	}

	mf.isEven = !mf.isEven
}

// FindMedian 给出 Median
func (mf *MedianFinder) FindMedian() float64 {
	if mf.isEven {
		return float64(mf.left.myHeap[0]+mf.right.myHeap[0]) / 2
	}
	return float64(mf.left.myHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type leftHeap struct {
	myHeap
}

func (l leftHeap) Less(i, j int) bool {
	return l.myHeap[i] > l.myHeap[j]
}

type rightHeap struct {
	myHeap
}

func (r rightHeap) Less(i, j int) bool {
	return r.myHeap[i] < r.myHeap[j]
}

// myHeap 实现了 heap 的接口
type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
