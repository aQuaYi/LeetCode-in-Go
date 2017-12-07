package Problem0539

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	n := len(timePoints)

	minDiff := 24*60 - diff(timePoints[0], timePoints[n-1])

	for i := 1; i < n; i++ {
		minDiff = min(minDiff, diff(timePoints[i-1], timePoints[i]))
	}

	return minDiff
}

// 返回 从 a 到 b 的分钟数
func diff(a, b string) int {
	ha, ma := analyze(a)
	hb, mb := analyze(b)

	return (hb-ha)*60 + (mb - ma)
}

func analyze(s string) (hour, minute int) {
	hour, _ = strconv.Atoi(s[:2])
	minute, _ = strconv.Atoi(s[3:])
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
