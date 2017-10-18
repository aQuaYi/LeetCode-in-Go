package Problem0630

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	taken := new(maxPQ)
	heap.Init(taken)

	css := cs(courses)
	sort.Sort(css)

	var date int
	for _, c := range css {
		heap.Push(taken, c[0])
		date += c[0]
		for date > c[1] {
			date -= heap.Pop(taken).(int)
		}
	}

	return taken.Len()
}

type maxPQ []int

// Len is
func (q maxPQ) Len() int {
	return len(q)
}

func (q maxPQ) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q maxPQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *maxPQ) Push(x interface{}) {
	// Push 使用 *q，是因为
	// Push 增加了 q 的长度
	*q = append(*q, x.(int))
}

func (q *maxPQ) Pop() interface{} {
	// Pop 使用 *q ，是因为
	// Pop 减短了 q 的长度
	res := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return res
}

type cs [][]int

func (c cs) Len() int {
	return len(c)
}

func (c cs) Less(i, j int) bool {
	if c[i][1] == c[j][1] {
		return c[i][0] < c[j][0]
	}
	return c[i][1] < c[j][1]
}

func (c cs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
