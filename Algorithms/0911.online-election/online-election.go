package problem0911

import "container/heap"

// TopVotedCandidate object will be instantiated and called as such:
// obj := Constructor(persons, times);
// param_1 := obj.Q(t);
type TopVotedCandidate struct {
}

// Constructor returns TopVotedCandidate
func Constructor(persons []int, times []int) TopVotedCandidate {

	return TopVotedCandidate{}
}

// Q is question
func (tvc *TopVotedCandidate) Q(t int) int {

	return 0
}

// entry 是 priorityQueue 中的元素
type entry struct {
	votedNumber int
	lastIndex   int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].lastIndex < pq[j].lastIndex
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the priority and value of an entry in the queue.
func (pq *PQ) update(entry *entry, value, priority int) {
	entry.votedNumber = value
	entry.lastIndex = priority
	heap.Fix(pq, entry.index)
}
