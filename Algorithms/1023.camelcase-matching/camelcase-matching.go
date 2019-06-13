package problem1023

func camelMatch(queries []string, pattern string) []bool {
	pr, pn := []rune(pattern), len(pattern)

	check := func(q string) bool {
		pi := 0
		for _, r := range q {
			if pi < pn && pr[pi] == r {
				pi++
				continue
			}
			if 'A' <= r && r <= 'Z' {
				return false
			}
		}
		return pi == pn
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = check(q)
	}
	return res
}
