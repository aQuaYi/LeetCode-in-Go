package kit

// This example demonstrates a priority queue built using the heap interface.

import (
	"container/heap"
)

// item 是 priorityQueue 中的元素
type item struct {
	value    string
	priority int
	// index 是 item 在 heap 中的索引号
	// item 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 item.priority 一直不变的话，可以删除 index
	index int
}

// priorityQueue implements heap.Interface and holds items.
type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 item
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop 从 pq 中取出最优先的 item
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an item in the queue.
func (pq *priorityQueue) update(item *item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
