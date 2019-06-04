package problem1001

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	l := newLights(N)
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		l.turnOn(x, y)
	}

	res := make([]int, 0, len(queries))
	for _, cell := range queries {
		x, y := cell[0], cell[1]
		res = append(res, l.checkIlluminated(x, y))
	}

	return res
}

type lights struct {
	h, v, ad, dd map[int]int
	isOn         map[int]bool
	n            int
}

func newLights(N int) *lights {
	return &lights{
		h:    make(map[int]int, 1024),
		v:    make(map[int]int, 1024),
		ad:   make(map[int]int, 1024),
		dd:   make(map[int]int, 1024),
		isOn: make(map[int]bool, 1024),
		n:    N,
	}
}

func (l *lights) turnOn(x, y int) {
	l.h[x]++
	l.v[y]++
	l.ad[x+y]++
	l.dd[x-y]++
	l.isOn[x<<32+y] = true
}

func (l *lights) turnOff(x, y int) {
	if x < 0 || l.n <= x ||
		y < 0 || l.n <= y ||
		!l.isOn[x<<32+y] {
		return
	}
	l.isOn[x<<32+y] = false
	l.h[x]--
	l.v[y]--
	l.ad[x+y]--
	l.dd[x-y]--
}

func (l *lights) checkIlluminated(x, y int) int {
	res := 0

	if l.h[x] > 0 { // maybe [x,y-1] or [x,y] or [x,y+1] is On
		res = 1
		l.turnOff(x, y-1)
		l.turnOff(x, y+1)
		// leave [x,y] in the end
	}

	if l.v[y] > 0 {
		res = 1
		l.turnOff(x-1, y)
		l.turnOff(x+1, y)
	}

	if l.ad[x+y] > 0 {
		res = 1
		l.turnOff(x-1, y+1)
		l.turnOff(x+1, y-1)
	}

	if l.dd[x-y] > 0 {
		res = 1
		l.turnOff(x-1, y-1)
		l.turnOff(x+1, y+1)
	}

	if res == 1 {
		l.turnOff(x, y)
	}

	return res
}
