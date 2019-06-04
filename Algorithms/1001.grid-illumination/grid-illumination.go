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
		if !l.isIlluminated(x, y) {
			res = append(res, 0)
			continue
		}
		res = append(res, 1)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				l.turnOff(x+dx, y+dy)
			}
		}
	}

	return res
}

type lights struct {
	h, v, ad, dd map[int]int
	hasLamp      map[int]bool
}

func newLights(N int) *lights {
	return &lights{
		h:       make(map[int]int, 1024),
		v:       make(map[int]int, 1024),
		ad:      make(map[int]int, 1024),
		dd:      make(map[int]int, 1024),
		hasLamp: make(map[int]bool, 1024),
	}
}

func (l *lights) turnOn(x, y int) {
	l.h[x]++
	l.v[y]++
	l.ad[x+y]++
	l.dd[x-y]++
	l.hasLamp[x<<32+y] = true
}

func (l *lights) turnOff(x, y int) {
	if !l.hasLamp[x<<32+y] {
		return
	}
	l.hasLamp[x<<32+y] = false
	l.h[x]--
	l.v[y]--
	l.ad[x+y]--
	l.dd[x-y]--
}

func (l *lights) isIlluminated(x, y int) bool {
	return l.h[x] > 0 ||
		l.v[y] > 0 ||
		l.ad[x+y] > 0 ||
		l.dd[x-y] > 0
}
