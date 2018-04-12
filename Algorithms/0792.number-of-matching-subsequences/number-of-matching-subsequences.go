package problem0792

func numMatchingSubseq(s string, words []string) int {
	dic := make(map[string]int, len(words))
	for _, w := range words {
		dic[w]++
	}

	res := 0
	for w := range dic {
		if isMatching(s, w) {
			res += dic[w]
		}
	}

	return res
}

func isMatching(s, w string) bool {
	m, n, j := len(s), len(w), 0
	for i := 0; i < m && j < n; i++ {
		if s[i] == w[j] {
			j++
		}
	}
	return j == n
}
