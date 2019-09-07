package problem1178

import "strings"

func findNumOfValidWords(words []string, puzzles []string) []int {
	m := len(words)
	lss := make([][]string, m)
	for i, w := range words {
		lss[i] = unique(w)
	}

	n := len(puzzles)
	res := make([]int, n)
	for i, p := range puzzles {
		for j, w := range words {
			res[i] += valid(p, w, lss[j])
		}
	}

	return res
}

func unique(w string) []string {
	n := len(w)
	lm := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		lm[w[i:i+1]] = true
	}
	ls := make([]string, 0, n)
	for l := range lm {
		ls = append(ls, l)
	}
	return ls
}

func valid(p, w string, ls []string) int {
	if !strings.Contains(w, p[:1]) {
		return 0
	}
	for _, l := range ls {
		if !strings.Contains(p, l) {
			return 0
		}
	}
	return 1
}
