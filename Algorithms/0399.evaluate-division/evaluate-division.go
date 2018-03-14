package problem0399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 建立变量之间的转换关系
	m := make(map[string]map[string]float64)
	for i, e := range equations {
		a, b := e[0], e[1]
		v := values[i]
		// 添加 a / b 的记录
		if _, ok := m[a]; !ok {
			m[a] = make(map[string]float64)
		}
		m[a][b] = 1.0 / v
		// 添加 b / a 的记录
		if _, ok := m[b]; !ok {
			m[b] = make(map[string]float64)
		}
		m[b][a] = v
	}

	// 逐个搜索 queries 的结果
	res := make([]float64, len(queries))
	for i, q := range queries {
		res[i] = bfs(m, q[0], q[1])
	}

	return res
}

type entry struct {
	s string
	f float64
}

func bfs(m map[string]map[string]float64, a, b string) float64 {
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

	isVisited := make(map[string]bool)
	queue := []entry{{a, 1.0}}

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if e.s == b {
			// 找到了 b
			return 1.0 / e.f
		}

		if isVisited[e.s] {
			continue
		}

		isVisited[e.s] = true

		for k, v := range m[e.s] {
			queue = append(queue, entry{k, v * e.f})
		}
	}
	//没有找到 b
	return -1.0
}
