package problem1023

import "strings"

func camelMatch(queries []string, pattern string) []bool {
	ps := split(pattern)
	res := make([]bool, len(queries))
	for qi, q := range queries {
		pi := 0
		i, j := 0, 0
		for pi < len(ps) {
			i, j = j, strings.Index(q, ps[pi])
			if i > j {
				break
			}
			q = strings.Replace(q, ps[pi], "", 1)
			pi++
		}
		res[qi] = pi == len(ps) && q == strings.ToLower(q)
	}
	return res
}

func split(pattern string) []string {
	var sb strings.Builder
	res := make([]string, 0, len(pattern)+1)
	for _, r := range pattern {
		if 'A' <= r && r <= 'Z' {
			res = append(res, sb.String())
			sb.Reset()
		}
		sb.WriteRune(r)
	}
	res = append(res, sb.String())
	return res[1:]
}
