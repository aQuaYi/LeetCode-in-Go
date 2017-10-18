package Problem0324

import "container/heap"

func wiggleSort(nums []int) {
	n := len(nums)

	// 寻找 中位数 mid
	left := new(maxHeap)
	left.intHeap = nums[:n/2]
	heap.Init(left)

	right := new(minHeap)
	right.intHeap = nums[n/2:]
	heap.Init(right)

	for left.intHeap[0] > right.intHeap[0] {
		left.intHeap[0], right.intHeap[0] = right.intHeap[0], left.intHeap[0]
		heap.Fix(left, 0)
		heap.Fix(right, 0)
	}

	mid := left.intHeap[0]

	// 3 路划分
	i, j, k := 0, 0, n-1
	for j <= k {
		if nums[j] > mid {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] < mid {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}
}

// maxHeap.intHeap[0] == max(maxHeap.intHeap)
type maxHeap struct {
	intHeap
}

func (h maxHeap) Less(i, j int) bool {
	return h.intHeap[i] > h.intHeap[j]
}

// minHeap.intHeap[0] == min(minHeap.intHeap)
type minHeap struct {
	intHeap
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
