package Problem0689

import "container/heap"
import "sort"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	res := make([]int, 0, 3)

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	pq := make(PQ, 0, len(nums)-k)
	pq = append(pq, &entry{starting: 0, sum: sum})

	for i := 1; i-1+k < len(nums); i++ {
		sum = sum - nums[i-1] + nums[i-1+k]
		pq = append(pq, &entry{starting: i, sum: sum})
	}

	heap.Init(&pq)

	for len(res) < 3 {
		temp := heap.Pop(&pq).(*entry).starting
		if !isOverlapping(res, temp, k) {
			res = append(res, temp)
		}
	}

	sort.Ints(res)

	return res
}

func isOverlapping(nums []int, x, k int) bool {
	for _, n := range nums {
		if isOverlap(n, x, k) {
			return true
		}
	}
	return false
}

func isOverlap(x, y, k int) bool {
	if x > y {
		x, y = y, x
	}
	return x <= y && y < x+k
}

// entry 是 priorityQueue 中的元素
type entry struct {
	starting, sum int
}

// PQ implements heap.Interface and holds entries.
type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].sum == pq[j].sum {
		return pq[i].starting < pq[j].starting
	}
	return pq[i].sum > pq[j].sum
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
