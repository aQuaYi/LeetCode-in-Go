package problem0720

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool, len(words))

	res := words[0]

	for _, w := range words {
		n := len(w)
		if n == 1 {
			m[w] = true
		} else if m[w[:n-1]] {
			m[w] = true
			if len(res) < len(w) {
				res = w
			}
		}
	}

	return res
}
