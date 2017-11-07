package Problem0460

import "container/heap"

// LFUCache 实现了 Least Frequently Used (LFU) cache
type LFUCache struct {
}

// Constructor 构建了 LFUCache
func Constructor(capacity int) LFUCache {

	return LFUCache{}
}

// Get 获取 key 的值
func (c *LFUCache) Get(key int) int {

	return 0
}

// Put 把 key， value 成对地放入 LFUCache
// 如果 LFUCache 已满的话，会删除掉 LFUCache 中使用最少的 key
func (c *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

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
