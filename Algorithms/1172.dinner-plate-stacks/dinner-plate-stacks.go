package problem1172

// DinnerPlates is ...
type DinnerPlates struct {
	cap int
	// plate's index of next push
	in int
	// plate's index of next pop
	// when out is -1, plates is empty, unable to pop
	out    int
	plates []*plate
}

// Constructor is ...
func Constructor(capacity int) DinnerPlates {
	plates := make([]*plate, 1, 1024)
	plates[0] = newPlate(capacity)
	return DinnerPlates{
		cap:    capacity,
		in:     0,
		out:    -1, // not possible to pop at beginning
		plates: plates,
	}
}

// Push is ...
func (d *DinnerPlates) Push(val int) {
	d.plates[d.in].push(val)
	// after push into a empty plate at end
	// d.out need point to the last nonempty plate
	if d.out < d.in {
		d.out = d.in
	}
	// make d.in to be the index of left-most nonfull plate
	for d.in < len(d.plates) && d.plates[d.in].isFull() {
		d.in++
	}
	// if no nonfull plate , create a new plate
	// JUST NOW, d.out < d.in
	if d.in == len(d.plates) {
		d.plates = append(d.plates, newPlate(d.cap))
	}
}

// Pop is a special condition of PopAtStack
func (d *DinnerPlates) Pop() int {
	if d.out == -1 {
		return -1
	}
	return d.PopAtStack(d.out)
}

// PopAtStack is ...
func (d *DinnerPlates) PopAtStack(i int) (res int) {
	if len(d.plates) <= i {
		return -1
	}
	p := d.plates[i]
	// set value and remove it from the plate
	if p.isEmpty() {
		return -1
	}
	res = p.pop()
	// make d.in to be the index of left-most nonfull plate
	d.in = min(d.in, i)
	// PopAtStack could make some empty plate in d.plates
	// need jump over these holes
	// make sure d.plates[d.out] have val to pop
	for d.out >= 0 && d.plates[d.out].isEmpty() {
		d.out--
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type plate struct {
	cap   int
	stack []int
}

func newPlate(cap int) *plate {
	return &plate{
		cap:   cap,
		stack: make([]int, 0, cap),
	}
}

func (p *plate) push(val int) {
	p.stack = append(p.stack, val)
}

func (p *plate) pop() (res int) {
	n := len(p.stack)
	p.stack, res = p.stack[:n-1], p.stack[n-1]
	return res
}

func (p *plate) isEmpty() bool {
	return len(p.stack) == 0
}

func (p *plate) isFull() bool {
	return len(p.stack) == p.cap
}
