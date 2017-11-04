package kit

// This example demonstrates a priority queue built using the heap interface.

import (
	"container/heap"
)

// entry 是 priorityQueue 中的元素
type entry struct {
	key      string
	priority int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*entry)
	entry.index = n
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1 // for safety
	*pq = old[0 : n-1]
	return entry
}

// update modifies the priority and value of an entry in the queue.
func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
