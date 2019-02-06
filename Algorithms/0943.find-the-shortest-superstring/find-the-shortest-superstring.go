package problem0943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	isUsed := make([]bool, size)
	suffixes := getSuffixes(A)
	res := strings.Repeat("?", 12*20+1)
	for i := 0; i < size; i++ {
		isUsed[i] = true
		greedy(A[i], i, size-1, A, isUsed, suffixes, &res)
		isUsed[i] = false
	}
	return res
}

func greedy(tmp string, endIndex, countDown int, A []string, isUsed []bool, suffixes [][]int, minRes *string) {
	if countDown == 0 {
		if len(*minRes) > len(tmp) {
			*minRes = tmp
		}
		return
	}

	// get max suffix length of UNUSED string
	maxLen := -1
	lens := suffixes[endIndex]
	for i, sl := range lens {
		if isUsed[i] {
			continue
		}
		maxLen = max(maxLen, sl)
	}

	// only connect string with max suffix length
	for i, sl := range lens {
		if sl < maxLen || isUsed[i] {
			continue
		}
		isUsed[i] = true
		greedy(tmp+A[i][sl:], i, countDown-1, A, isUsed, suffixes, minRes)
		isUsed[i] = false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// res[i][j] == 3 means A[j][:3] is A[i]'s suffix
func getSuffixes(A []string) [][]int {
	size := len(A)
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			res[i][j] = suffix(A[i], A[j])
		}
	}
	return res
}

func suffix(a, b string) int {
	i := len(b)
	for !strings.HasSuffix(a, b[:i]) {
		i--
	}
	return i
}
