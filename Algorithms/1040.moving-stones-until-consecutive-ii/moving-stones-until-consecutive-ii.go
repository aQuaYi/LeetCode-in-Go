package problem1040

import "sort"

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	return []int{minMove(stones), maxMove(stones)}
}

func maxMove(s []int) int {
	n := len(s)
	return max(s[n-1]-s[1], s[n-2]-s[0]) - n + 2
}

func minMove(s []int) int {
	n := len(s)
	width := 0
	i, j := 0, 1
	for ; j < n; j++ {
		if s[j]-s[i] < n {
			continue
		}
		width = max(width, j-i)
		i = j
	}
	width = max(width, j-i)
	return n - width
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
