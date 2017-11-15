package Problem0502

import "container/heap"

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	size := len(Profits)
	pros := make(proPQ, 0, size)
	caps := make(capPQ, size)

	for i := range Profits {
		p := &project{
			profit:  Profits[i],
			capital: Capital[i],
		}
		caps[i] = p
	}

	heap.Init(&caps)

	for {
		for len(caps) > 0 && caps[0].capital <= W {
			heap.Push(&pros, heap.Pop(&caps))
		}

		if len(pros) == 0 || k == 0 {
			break
		}

		W += heap.Pop(&pros).(*project).profit
		k--
	}

	return W
}

// entry 是 priorityQueue 中的元素
type project struct {
	profit, capital int
}

// PQ implements heap.Interface and holds entries.
type capPQ []*project

func (pq capPQ) Len() int { return len(pq) }

func (pq capPQ) Less(i, j int) bool {
	return pq[i].capital < pq[j].capital
}

func (pq capPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *capPQ) Push(x interface{}) {
	entry := x.(*project)
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *capPQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}

// PQ implements heap.Interface and holds entries.
type proPQ []*project

func (pq proPQ) Len() int { return len(pq) }

func (pq proPQ) Less(i, j int) bool {
	return pq[i].profit > pq[j].profit
}

func (pq proPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 entry
func (pq *proPQ) Push(x interface{}) {
	entry := x.(*project)
	*pq = append(*pq, entry)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *proPQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}
