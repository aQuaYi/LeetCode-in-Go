package problem0792

import (
	"strings"
)

func numMatchingSubseq(S string, words []string) int {
	res := 0
	for _, w := range words {
		if isMatching(S, w) {
			res++
		}
	}
	return res
}

func isMatching(s, c string) bool {
	idx := 0

	for i := 0; i < len(c); i++ {
		tempIdx := strings.Index(s[idx:], c[i:i+1])
		if tempIdx == -1 {
			return false
		}
		idx += tempIdx + 1
	}

	return true
}
