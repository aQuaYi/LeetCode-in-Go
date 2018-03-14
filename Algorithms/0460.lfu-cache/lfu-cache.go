package problem0460

import "container/heap"
import "time"

// LFUCache 实现了 Least Frequently Used (LFU) cache
type LFUCache struct {
	// 用于检查 key 的存在性
	m   map[int]*entry
	pq  PQ
	cap int
}

// Constructor 构建了 LFUCache
func Constructor(capacity int) LFUCache {
	return LFUCache{
		m:   make(map[int]*entry, capacity),
		pq:  make(PQ, 0, capacity),
		cap: capacity,
	}
}

// Get 获取 key 的值
func (c *LFUCache) Get(key int) int {
	ep, ok := c.m[key]
	if ok {
		c.pq.update(ep)
		return ep.value
	}
	return -1
}

// Put 把 key， value 成对地放入 LFUCache
// 如果 LFUCache 已满的话，会删除掉 LFUCache 中使用最少的 key
func (c *LFUCache) Put(key int, value int) {
	if c.cap <= 0 {
		return
	}
	ep, ok := c.m[key]
	// key 已存在，就更新 value
	if ok {
		c.m[key].value = value
		c.pq.update(ep)
		return
	}

	ep = &entry{key: key, value: value}
	// pq 已满的话，需要先删除，再插入
	if len(c.pq) == c.cap {
		temp := heap.Pop(&c.pq).(*entry)
		delete(c.m, temp.key)
	}
	c.m[key] = ep
	heap.Push(&c.pq, ep)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// entry 是 priorityQueue 中的元素
type entry struct {
	key, value int
	// 以下是辅助参数，由 heap.Interface 实现的函数自动处理
	frequence, index int
	date             time.Time
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].frequence == pq[j].frequence {
		return pq[i].date.Before(pq[j].date)
	}

	return pq[i].frequence < pq[j].frequence
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
	entry.date = time.Now()
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

// update modifies the priority of an entry in the queue.
func (pq *PQ) update(entry *entry) {
	entry.frequence++
	entry.date = time.Now()
	heap.Fix(pq, entry.index)
}
