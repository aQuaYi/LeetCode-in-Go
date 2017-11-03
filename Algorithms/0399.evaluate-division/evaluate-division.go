package Problem0399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 建立变量之间的转换关系
	m := make(map[string]map[string]float64)
	for i, e := range equations {
		a, b := e[0], e[1]
		v := values[i]
		if _, ok := m[a]; !ok {
			m[a] = make(map[string]float64)
		}
		m[a][b] = 1.0 / v
		if _, ok := m[b]; !ok {
			m[b] = make(map[string]float64)
		}
		m[b][a] = v
	}

	// 逐个查询 queries 的结果
	res := make([]float64, 0, len(queries))
	for _, q := range queries {
		res = append(res, query(m, q[0], q[1]))
	}

	return res
}

type entry struct {
	s string
	f float64
}

func query(m map[string]map[string]float64, a, b string) float64 {
	_, ok := m[a]
	if !ok {
		return -1.0
	}
	_, ok = m[b]
	if !ok {
		return -1.0
	}

	if a == b {
		return 1.0
	}

	visited := make(map[string]bool)
	queue := []entry{{a, 1.0}}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.s == b {
			return 1.0 / top.f
		}
		if visited[top.s] {
			continue
		}
		visited[top.s] = true
		for k, v := range m[top.s] {
			queue = append(queue, entry{k, v * top.f})
		}
	}
	return -1.0
}
