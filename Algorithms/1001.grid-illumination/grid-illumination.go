package problem1001

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	l := newLights(N)
	for _, lamp := range lamps {
		l.turnOn(lamp)
	}

	res := make([]int, 0, len(queries))
	for _, cell := range queries {
		res = append(res, l.illuminated(cell))
		x, y := cell[0], cell[1]
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				l.turnOff([]int{x + i, y + j})
			}
		}
	}

	return res
}

type lights struct {
	h, v         map[int]int
	ld2ru, lu2rd map[[2]int]int
	hasLamp      map[[2]int]bool
	N            int
}

func newLights(N int) *lights {
	return &lights{
		h:       make(map[int]int, N),
		v:       make(map[int]int, N),
		ld2ru:   make(map[[2]int]int, N*N),
		lu2rd:   make(map[[2]int]int, N*N),
		hasLamp: make(map[[2]int]bool, N*N),
		N:       N,
	}
}

func (l *lights) turnOn(lamp []int) {
	x, y := lamp[0], lamp[1]
	l.h[x]++
	l.v[y]++
	l.ld2ru[ld(lamp, l.N)]++
	l.lu2rd[lu(lamp, l.N)]++
	l.hasLamp[[2]int{x, y}] = true
}

func (l *lights) turnOff(lamp []int) {
	x, y := lamp[0], lamp[1]
	if !l.hasLamp[[2]int{x, y}] {
		return
	}
	delete(l.hasLamp, [2]int{x, y})
	l.h[x] = max(0, l.h[x]-1)
	l.v[y] = max(0, l.v[y]-1)
	ld2ru := ld(lamp, l.N)
	l.ld2ru[ld2ru] = max(0, l.ld2ru[ld2ru]-1)
	lu2rd := lu(lamp, l.N)
	l.lu2rd[lu2rd] = max(0, l.lu2rd[lu2rd]-1)
}

func (l *lights) illuminated(cell []int) int {
	x, y := cell[0], cell[1]
	if l.h[x] > 0 ||
		l.v[y] > 0 ||
		l.ld2ru[ld(cell, l.N)] > 0 ||
		l.lu2rd[lu(cell, l.N)] > 0 {
		return 1
	}
	return 0
}

// form lamp to its left-down point
func ld(lamp []int, N int) [2]int {
	x, y := lamp[0], lamp[1]
	d := min(x, N-1-y)
	return [2]int{x + d, y - d}
}

// form lamp to its left-up point
func lu(lamp []int, N int) [2]int {
	x, y := lamp[0], lamp[1]
	d := min(x, y)
	return [2]int{x - d, y - d}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
