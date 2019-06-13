package problem1023

import "strings"

func camelMatch(queries []string, pattern string) []bool {
	pr, pn := []rune(pattern), len(pattern)

	check := func(q string) bool {
		var sb strings.Builder
		pi := 0
		for _, r := range q {
			if pi < pn && pr[pi] == r {
				pi++
				continue
			}
			if 'A' <= r && r <= 'Z' {
				return false
			}
			sb.WriteRune(r)
		}
		q = sb.String()
		return pi == pn &&
			q == strings.ToLower(q)
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = check(q)
	}
	return res
}
