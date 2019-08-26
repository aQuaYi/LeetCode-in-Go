package problem1170

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	n := len(words)
	ws := make([]int, n)
	for i, w := range words {
		ws[i] = f(w)
	}
	sort.Ints(ws)

	res := make([]int, len(queries))
	for i, q := range queries {
		fq := f(q)
		res[i] = n - sort.Search(n, func(i int) bool { return fq < ws[i] })
	}

	return res
}

func f(s string) int {
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	i := 0
	for count[i] == 0 {
		i++
	}
	return count[i]
}
