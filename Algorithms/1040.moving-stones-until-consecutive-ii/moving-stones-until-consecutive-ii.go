package problem1040

import "sort"

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	return []int{low(stones), high(stones)}
}

func high(s []int) int {
	n := len(s)
	return max(s[n-1]-s[1], s[n-2]-s[0]) - n + 2
}

func low(s []int) int {
	n := len(s)
	// corner case
	if (s[n-2]-s[0] == n-2 && s[n-1]-s[n-2] > 2) ||
		(s[n-1]-s[1] == n-2 && s[1]-s[0] > 2) {
		return 2
	}
	// sliding window is s[i:j]
	width, i, j := 0, 0, 1
	for ; j < n; j++ {
		if s[j]-s[i] < n {
			continue
		}
		width = max(width, j-i)
		i = j
	}
	width = max(width, j-i)
	// finally, all stone move into maxWidth windows
	// so need move n-width stones
	return n - width
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
