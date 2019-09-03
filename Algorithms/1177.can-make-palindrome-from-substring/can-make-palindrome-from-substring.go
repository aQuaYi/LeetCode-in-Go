package problem1177

import "sort"

func canMakePaliQueries(s string, queries [][]int) []bool {
	indexesOf := [26][]int{}
	for i, letter := range s {
		b := letter - 'a'
		indexesOf[b] = append(indexesOf[b], i)
	}

	count := func(left, right int) int {
		single := 0
		for _, indexes := range indexesOf {
			l := sort.SearchInts(indexes, left)
			r := sort.SearchInts(indexes, right+1)
			single += (r - l) % 2
		}
		return single / 2
	}

	n := len(queries)
	needs := make(map[int]int, n)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		q := queries[i]
		left, right, k := q[0], q[1], q[2]
		m := mark(left, right)
		need, ok := needs[m]
		if !ok {
			need = count(left, right)
			needs[m] = need
		}
		res[i] = k >= need
	}

	return res
}

func mark(left, right int) int {
	return left<<18 + right
}
