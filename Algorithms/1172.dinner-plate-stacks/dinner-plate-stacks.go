package problem1172

import "container/heap"

// DinnerPlates is ...
type DinnerPlates struct {
	cap     int
	plates  PQ
	inOrder []*plate
}

func (d *DinnerPlates) isEmpty() bool {
	return len(d.plates) == 0
}

// Constructor is ...
func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		cap:     capacity,
		plates:  make(PQ, 0, 2048),
		inOrder: make([]*plate, 0, 2048),
	}
}

// Push is ...
func (d *DinnerPlates) Push(val int) {
	if d.isEmpty() || d.plates[0].isFull() {
		id := len(d.inOrder)
		p := newPlate(id, d.cap)
		p.push(val)
		heap.Push(&d.plates, p)
		d.inOrder = append(d.inOrder, p)
	} else {
		s := d.plates[0]
		s.push(val)
		heap.Fix(&d.plates, 0)
	}
}

// Pop is ...
func (d *DinnerPlates) Pop() int {
	if d.isEmpty() {
		return -1
	}
	n := len(d.inOrder)
	s := d.inOrder[n-1]
	if s.isEmpty() {
		d.inOrder = d.inOrder[:n-1]
		s.id = -1
		heap.Fix(&d.plates, s.index)
		heap.Pop(&d.plates)
		return d.Pop()
	}
	res := s.pop()
	return res
}

// PopAtStack is ...
func (d *DinnerPlates) PopAtStack(id int) int {
	if id > len(d.inOrder)-1 ||
		d.inOrder[id] == nil {
		return -1
	}
	s := d.inOrder[id]
	if s.isEmpty() {
		return -1
	}
	res := s.pop()
	heap.Fix(&d.plates, s.index)
	return res
}

// plate 是 priorityQueue 中的元素
type plate struct {
	id, cap, top, index int
	vals                []int
}

func newPlate(id, cap int) *plate {
	return &plate{
		id:   id,
		cap:  cap,
		top:  -1,
		vals: make([]int, cap),
	}
}

func (s *plate) isFull() bool {
	return s.top == s.cap-1
}

func (s *plate) isEmpty() bool {
	return s.top == -1
}

func (s *plate) pop() int {
	res := s.vals[s.top]
	s.top--
	return res
}

func (s *plate) push(num int) {
	s.top++
	s.vals[s.top] = num
}

// PQ implements heap.Interface and holds entries.
type PQ []*plate

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
	temp := x.(*plate)
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
