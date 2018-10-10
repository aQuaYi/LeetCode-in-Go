package problem0895

import (
	"container/heap"
)

// FreqStack object will be instantiated and called as such:
// obj := Constructor();
// obj.Push(x);
// param_2 := obj.Pop();
type FreqStack struct {
	index int
	freq  map[int]int
	pq    *PQ
}

// Constructor 构建 FreqStack
func Constructor() FreqStack {
	pq := make(PQ, 0, 10000)
	return FreqStack{
		freq: make(map[int]int, 10000),
		pq:   &pq,
	}
}

// Push 在 fs 中放入 x
func (fs *FreqStack) Push(x int) {
	fs.index++
	fs.freq[x]++
	e := &entry{
		key:   x,
		index: fs.index,
		freq:  fs.freq[x],
	}
	heap.Push(fs.pq, e)
}

// Pop 从 fs 中弹出元素
func (fs *FreqStack) Pop() int {
	x := heap.Pop(fs.pq).(*entry).key
	fs.freq[x]--
	return x
}

// entry 是 priorityQueue 中的元素
type entry struct {
	key, index, freq int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].freq == pq[j].freq {
		return pq[i].index > pq[j].index
	}
	return pq[i].freq > pq[j].freq
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
