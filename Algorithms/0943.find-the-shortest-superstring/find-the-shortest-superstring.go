package problem0943

import (
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	isUsed := make([]bool, size)
	overlaps := getOverlaps(A)
	res := strings.Repeat("?", 12*20+1)
	for i := 0; i < size; i++ {
		isUsed[i] = true
		greedy(A[i], i, size-1, A, isUsed, overlaps, &res)
		isUsed[i] = false
	}
	return res
}

func greedy(tmp string, last, countDown int, A []string, isUsed []bool, overlaps [][]int, res *string) {
	if countDown == 0 {
		if len(*res) > len(tmp) {
			*res = tmp
		}
		return
	}

	maxLen := -1
	lens := overlaps[last]

	for i, j := range lens {
		if isUsed[i] {
			continue
		}
		maxLen = max(maxLen, j)
	}

	for j, ol := range lens {
		if ol < maxLen || isUsed[j] {
			continue
		}
		isUsed[j] = true
		s := A[j]
		greedy(tmp+s[ol:], j, countDown-1, A, isUsed, overlaps, res)
		isUsed[j] = false
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// res[i][j] == 3 means A[j][:3] is A[i]'s suffix
func getOverlaps(A []string) [][]int {
	size := len(A)
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			res[i][j] = overlap(A[i], A[j])
		}
	}
	return res
}

func overlap(a, b string) int {
	i := len(b)
	for !strings.HasSuffix(a, b[:i]) {
		i--
	}
	return i
}
