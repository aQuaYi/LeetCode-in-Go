package problem0792

func numMatchingSubseq(S string, words []string) int {
	isChecked := make(map[string]bool, len(words))
	res := 0
	for _, w := range words {
		if isChecked[w] {
			res++
			continue
		}

		if isMatching(S, w) {
			isChecked[w] = true
			res++
		}
	}

	return res
}

func isMatching(s, c string) bool {
	m, n, j := len(s), len(c), 0
	for i := 0; i < m && j < n; i++ {
		if s[i] == c[j] {
			j++
		}
	}
	return j == n
}
