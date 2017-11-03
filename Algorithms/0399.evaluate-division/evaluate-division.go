package Problem0399

import "sort"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]float64, len(equations)*2)
	z := make(map[string]int, len(equations)*2)
	zi := 1

	ev := equationsAndValues{e: equations, v: values}

	sort.Sort(ev)

	for i := len(ev.e) - 1; 0 <= i; i-- {
		v1, ok1 := m[ev.e[i][1]]
		if ok1 {
			m[ev.e[i][0]] = v1 * ev.v[i]
			z[ev.e[i][0]] = zi
		} else {
			zi++
			z[ev.e[i][1]] = zi
			z[ev.e[i][0]] = zi
			m[ev.e[i][1]] = 1.0
			m[ev.e[i][0]] = ev.v[i]
		}
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		v0, ok0 := m[q[0]]
		v1, ok1 := m[q[1]]

		if ok0 && ok1 && z[q[0]] == z[q[1]] {
			res[i] = v0 / v1
		} else {
			res[i] = -1.0
		}
	}

	return res
}

type equationsAndValues struct {
	e [][]string
	v []float64
}

func (ev equationsAndValues) Len() int {
	return len(ev.e)
}

func (ev equationsAndValues) Less(i, j int) bool {
	return ev.e[i][0] < ev.e[j][0]
}

func (ev equationsAndValues) Swap(i, j int) {
	ev.e[i], ev.e[j] = ev.e[j], ev.e[i]
	ev.v[i], ev.v[j] = ev.v[j], ev.v[i]
}
