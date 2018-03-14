package problem0502

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	size := len(Profits)

	// caps 按照 capital 的升序排列 project
	caps := make(capQueue, size)
	for i := range Profits {
		p := &project{
			profit:  Profits[i],
			capital: Capital[i],
		}
		caps[i] = p
	}
	sort.Sort(caps)

	// pros 是高 profit 优先的队列
	pros := make(proPQ, 0, size)
	var i int
	for {
		// 把所有可以做的 project 存入 pros
		for i < len(caps) && caps[i].capital <= W {
			heap.Push(&pros, caps[i])
			i++
		}

		// 无 project 可做了
		// 或
		// 已经达到工作量的上限
		if len(pros) == 0 || k == 0 {
			break
		}

		// 从 pros 中挑选一件 profit 最大的工作来做
		// 并更新 W
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
type capQueue []*project

func (q capQueue) Len() int { return len(q) }

func (q capQueue) Less(i, j int) bool {
	return q[i].capital < q[j].capital
}

func (q capQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
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
