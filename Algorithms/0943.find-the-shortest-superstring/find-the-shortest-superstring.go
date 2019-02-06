package problem0943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	isUsed := make([]bool, size)
	res := strings.Repeat("?", 12*20+1)
	greedy("", size, A, isUsed, &res)
	return res
}

func greedy(tmp string, countDown int, A []string, isUsed []bool, res *string) {
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
		for !strings.HasSuffix(tmp, str[:j]) { // heavy operation
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
		greedy(tmp+s[j:], countDown-1, A, isUsed, res)
		isUsed[i] = false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
