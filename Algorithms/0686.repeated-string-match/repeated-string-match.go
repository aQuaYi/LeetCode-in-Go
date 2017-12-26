package Problem0686

import "strings"

func repeatedStringMatch(a string, b string) int {
	times := max(len(b)/len(a), 1)

	if strings.Contains(strings.Repeat(a, times), b) {
		return times
	}
	if strings.Contains(strings.Repeat(a, times+1), b) {
		return times + 1
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
