package problem1041

func isRobotBounded(instructions string) bool {
	print(north, east, south, west)
	r := newRobot()
	for _, i := range instructions {
		r.receive(i)
	}
	return (r.x == 0 && r.y == 0) || // 每轮都会回到原点，就永远会回到原点
		(r.d != north) // 第一轮后，没有回到原点，但是方向改变，至多 4 轮后，还是会回到原点
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
		r.d += 3
	}
	r.d %= 4
}
