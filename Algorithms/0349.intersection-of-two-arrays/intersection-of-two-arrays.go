package Problem0349

func intersection(a1, a2 []int) []int {
	res := []int{}
	m1 := makeMap(a1)
	m2 := makeMap(a2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	for n := range m1 {
		if _, ok := m2[n]; ok {
			res = append(res, n)
		}
	}

	return res
}

func makeMap(a []int) map[int]bool {
	res := make(map[int]bool, len(a))

	for i := range a {
		res[a[i]] = true
	}

	return res
}
