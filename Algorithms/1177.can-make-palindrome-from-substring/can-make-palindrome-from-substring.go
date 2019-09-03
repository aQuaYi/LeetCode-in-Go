package problem1177

import "sort"

func canMakePaliQueries(s string, queries [][]int) []bool {
	count := [26][]int{}
	for i, letter := range s {
		b := letter - 'a'
		count[b] = append(count[b], i)
	}

	check := func(q []int) bool {
		left, right, k := q[0], q[1]+1, q[2]+1
		hasSeen := [26]bool{}
		for i := left; i < right; i++ {
			b := s[i] - 'a'
			if hasSeen[b] {
				continue
			}
			hasSeen[b] = true
			cs := count[b]
			l := sort.SearchInts(cs, left)
			r := sort.SearchInts(cs, right)
			k -= (r - l) % 2
			if k < 0 {
				return false
			}
		}
		return true
	}
	n := len(queries)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = check(queries[i])
	}

	return res
}
