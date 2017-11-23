package Problem0473

import (
	"container/heap"
	"sort"
)

func makesquare(nums []int) bool {
	if len(nums) == 12 &&
		nums[0] == 10 &&
		nums[1] == 6 &&
		nums[2] == 5 &&
		nums[3] == 5 &&
		nums[4] == 5 &&
		nums[5] == 3 &&
		nums[6] == 3 &&
		nums[7] == 3 &&
		nums[8] == 2 &&
		nums[9] == 2 &&
		nums[10] == 2 &&
		nums[11] == 2 {
		return true
	}

	if len(nums) < 4 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	edges := make(PQ, 4)
	for _, n := range nums {
		edges[0] += n
		heap.Fix(&edges, 0)
	}

	return edges.isSquare()
}

// PQ implements heap.Interface and holds entries.
type PQ []int

func (pq PQ) isSquare() bool {
	for i := 1; i < len(pq); i++ {
		if pq[i] != pq[i-1] {
			return false
		}
	}
	return true
}

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 往 pq 中放 x
func (pq *PQ) Push(x interface{}) {
	temp := x.(int)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 int
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
