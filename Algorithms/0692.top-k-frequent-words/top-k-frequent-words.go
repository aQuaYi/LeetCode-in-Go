package Problem0692

import "container/heap"

func topKFrequent(words []string, k int) []string {
	count := make(map[string]int, len(words))
	for _, w := range words {
		count[w]++
	}

	pq := make(PQ, 0, len(count))
	for w, c := range count {
		pq = append(pq, &entry{
			word:      w,
			frequence: c,
		})
	}
	heap.Init(&pq)

	res := make([]string, 0, k)
	for k > 0 {
		res = append(res, heap.Pop(&pq).(*entry).word)
		k--
	}

	return res
}

// entry 是 priorityQueue 中的元素
type entry struct {
	word      string
	frequence int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].frequence == pq[j].frequence {
		return pq[i].word < pq[j].word
	}
	return pq[i].frequence > pq[j].frequence
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
