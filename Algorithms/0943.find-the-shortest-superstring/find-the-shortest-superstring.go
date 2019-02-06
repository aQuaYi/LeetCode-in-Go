package problem0943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	res := strings.Repeat("?", 12*20+1)
	isUsed := make([]bool, size)
	rescur(A, isUsed, size, "", &res)
	return res
}

func rescur(A []string, isUsed []bool, countDown int, tmp string, res *string) {
	if countDown == 0 {
		if len(*res) > len(tmp) {
			*res = tmp
		}
		return
	}

	maxLen := -1
	lens := make([]int, len(A))

	for i, str := range A {
		if isUsed[i] {
			continue
		}
		j := len(str)
		for !strings.HasSuffix(tmp, str[:j]) {
			j--
		}
		lens[i] = j
		maxLen = max(maxLen, j)
	}

	for i, j := range lens {
		if j < maxLen || isUsed[i] {
			continue
		}
		isUsed[i] = true
		s := A[i]
		rescur(A, isUsed, countDown-1, tmp+s[j:], res)
		isUsed[i] = false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
