package problem1041

func isRobotBounded(instructions string) bool {
	r := newRobot()
	for _, i := range instructions {
		r.receive(i)
	}
	return (r.x == 0 && r.y == 0) ||
		(r.d != north)
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

type robot struct {
	d    direction
	x, y int
}

func newRobot() *robot {
	return &robot{
		d: north,
	}
}

func (r *robot) receive(instruction rune) {
	if instruction == 'G' {
		r.move()
	} else {
		r.turn(instruction)
	}
}

func (r *robot) move() {
	switch r.d {
	case north:
		r.y++
	case east:
		r.x++
	case south:
		r.y--
	case west:
		r.x--
	}
}

func (r *robot) turn(instruction rune) {
	if instruction == 'R' {
		r.d++
	} else {
		r.d--
	}
	r.d = r.d % 4
}
