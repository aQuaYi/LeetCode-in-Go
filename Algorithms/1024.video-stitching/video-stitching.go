package problem1024

import "sort"

func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i int, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	res := 0

	for i, start, end := 0, 0, 0; start < T; start = end {
		for i < len(clips) && clips[i][0] <= start {
			end = max(end, clips[i][1])
			i++
		}
		if start == end {
			return -1
		}
		res++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
