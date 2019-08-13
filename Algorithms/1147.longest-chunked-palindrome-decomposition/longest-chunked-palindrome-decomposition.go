package problem1147

import "strings"

func longestDecomposition(S string) int {
	if S == "" {
		return 0
	}
	n := len(S)
	for i := 1; i*2 <= n; i++ {
		left := S[:i]
		if strings.HasSuffix(S, left) {
			return 2 + longestDecomposition(S[i:n-i])
		}
	}
	return 1
}
