package Problem0399

import "sort"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]float64, len(equations)*2)

	sort.Sort(equationSlice(equations))

	for i, e := range equations {
		v0, ok0 := m[e[0]]
		v1, ok1 := m[e[1]]

		if ok0 && ok1 {
			continue
		}

		if ok0 {
			m[e[1]] = v0 / values[i]
			continue
		}

		if ok1 {
			m[e[0]] = v1 * values[i]
			continue
		}

		m[e[1]] = 1.0
		m[e[0]] = values[i]

	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		v0, ok0 := m[q[0]]
		v1, ok1 := m[q[1]]

		if ok0 && ok1 {
			res[i] = v0 / v1
		} else {
			res[i] = -1.0
		}
	}

	return res
}

type equationSlice [][]string

func (e equationSlice) Len() int {
	return len(e)
}

func (e equationSlice) Less(i, j int) bool {
	return e[i][0] < e[j][0]
}

func (e equationSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
