package problem0480

import "container/heap"
import "time"

func medianSlidingWindow(nums []int, k int) []float64 {
	res := make([]float64, len(nums)-k+1)
	w := newWindow(nums, k)

	for i := 0; i+k < len(nums); i++ {
		res[i] = w.median()
		w.update(i, i+k, nums)
	}

	res[len(nums)-k] = w.median()

	return res
}

type window struct {
	k     int
	medX2 int
	m     map[int]*entry
	max   maxPQ
	min   minPQ
}

func newWindow(nums []int, k int) window {
	max := make(maxPQ, 0, k)
	min := make(minPQ, 0, k)
	m := make(map[int]*entry, len(nums))

	for i := 0; i < k; i++ {
		ep := &entry{
			idx: i,
			val: nums[i],
		}

		m[i] = ep

		if min.Len() == max.Len() {
			heap.Push(&min, ep)
		} else {
			heap.Push(&max, ep)
		}
	}

	for len(max) > 0 && max[0].val > min[0].val {
		max[0], min[0] = min[0], max[0]
		heap.Fix(&max, 0)
		heap.Fix(&min, 0)
	}

	return window{
		k:   k,
		m:   m,
		max: max,
		min: min,
	}
}

func (w window) median() float64 {
	if w.k&1 == 1 {
		return float64(w.min[0].val)
	}
	return float64(w.max[0].val+w.min[0].val) / 2
}

func (w *window) update(popIdx, pushIdx int, nums []int) {
	ep := w.m[popIdx]

	ep.idx = pushIdx
	ep.val = nums[pushIdx]
	w.m[pushIdx] = ep

	if ep.index < len(w.max) {
		heap.Fix(&w.max, ep.index)
	}

	if ep.index < len(w.min) {
		heap.Fix(&w.min, ep.index)
	}

	time.Sleep(time.Millisecond * 10)
	if len(w.max) > 0 && w.max[0].val > w.min[0].val {
		w.max[0], w.min[0] = w.min[0], w.max[0]
		heap.Fix(&w.max, 0)
		heap.Fix(&w.min, 0)
	}
}

// entry 是 priorityQueue 中的元素
type entry struct {
	idx int
	val int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// minPQ implements heap.Interface and holds entries.
type minPQ []*entry

func (pq minPQ) Len() int { return len(pq) }

func (pq minPQ) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *minPQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*entry)
	entry.index = n
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1 // for safety
	*pq = old[0 : n-1]
	return entry
}

// maxPQ implements heap.Interface and holds entries.
type maxPQ []*entry

func (pq maxPQ) Len() int { return len(pq) }

func (pq maxPQ) Less(i, j int) bool {
	return pq[i].val > pq[j].val
}

func (pq maxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *maxPQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*entry)
	entry.index = n
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *maxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1 // for safety
	*pq = old[0 : n-1]
	return entry
}
