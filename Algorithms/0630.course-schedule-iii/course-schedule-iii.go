package problem0630

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	// taken 已经上过的课程的时长队列
	taken := new(maxPQ)
	heap.Init(taken)

	// myCs 具体课程表
	myCs := cs(courses)
	sort.Sort(myCs)

	// date 已经上过的课程的总时长
	var date int
	// for 循环始终保持以下最优状态
	// 1. c[1] 日期前，尽可能多的课程
	// 2. 如果必须舍弃一部分课程的话。舍弃时长最长的课程，让 date 最小。只有这样的话，才能让后面的有机会上更多的课程。
	for _, c := range myCs {
		heap.Push(taken, c[0])
		date += c[0]
		for date > c[1] {
			date -= heap.Pop(taken).(int)
		}
	}

	return taken.Len()
}

// maxPQ 是最大值优先的队列
type maxPQ []int

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

// cs 是课程的描述课程的数据结构
// cs[i][0]: 第 i 个课程的持续时间
// cs[i][1]: 第 i 个课程的结束截止日期
type cs [][]int

func (c cs) Len() int {
	return len(c)
}

func (c cs) Less(i, j int) bool {
	if c[i][1] == c[j][1] {
		// 相同截止日期时，持续时间短的课程靠前
		return c[i][0] < c[j][0]
	}
	// 结束时间早的靠前
	return c[i][1] < c[j][1]
}

func (c cs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
