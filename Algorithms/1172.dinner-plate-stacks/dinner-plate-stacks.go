package problem1172

import "container/heap"

// DinnerPlates is ...
type DinnerPlates struct {
	cap    int
	plates *PQ
	rec    []*stack
}

func (d *DinnerPlates) isEmpty() bool {
	return len(*d.plates) == 0
}

// Constructor is ...
func Constructor(capacity int) DinnerPlates {
	pq := make(PQ, 0, 1024)
	return DinnerPlates{
		cap:    capacity,
		plates: &pq,
		rec:    make([]*stack, 0, 1024),
	}
}

// Push is ...
func (d *DinnerPlates) Push(val int) {
	if d.isEmpty() || (*d.plates)[0].isFull() {
		id := len(d.rec)
		s := newStack(id, d.cap)
		s.push(val)
		heap.Push(d.plates, s)
		d.rec = append(d.rec, s)
	} else {
		s := (*d.plates)[0]
		s.push(val)
		heap.Fix(d.plates, 0)
	}
}

// Pop is ...
func (d *DinnerPlates) Pop() int {
	if d.isEmpty() {
		return -1
	}
	n := len(d.rec)
	s := d.rec[n-1]
	res := s.pop()
	if s.isEmpty() {
		d.rec = d.rec[:n-1]
		s.id = -1
		heap.Fix(d.plates, s.index)
		heap.Pop(d.plates)
	}
	return res
}

// PopAtStack is ...
func (d *DinnerPlates) PopAtStack(id int) int {
	if d.rec[id] == nil {
		return -1
	}
	s := d.rec[id]
	if s.isEmpty() {
		return -1
	}
	res := s.pop()
	heap.Fix(d.plates, s.index)
	return res
}

// stack 是 priorityQueue 中的元素
type stack struct {
	id, cap int
	vals    []int
	index   int
}

func newStack(id, cap int) *stack {
	return &stack{
		id:   id,
		cap:  cap,
		vals: make([]int, 0, cap),
	}
}

func (s *stack) isFull() bool {
	return len(s.vals) == s.cap
}

func (s *stack) isEmpty() bool {
	return len(s.vals) == 0
}

func (s *stack) pop() (top int) {
	n := len(s.vals)
	s.vals, top = s.vals[:n-1], s.vals[n-1]
	return top
}

func (s *stack) push(num int) {
	s.vals = append(s.vals, num)
}

// PQ implements heap.Interface and holds entries.
type PQ []*stack

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	switch {
	case pq[i].isFull() && pq[j].isFull():
		return pq[i].id > pq[j].id
	case pq[i].isFull():
		return false
	case pq[j].isFull():
		return true
	default:
		return pq[i].id < pq[j].id
	}
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*stack)
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
