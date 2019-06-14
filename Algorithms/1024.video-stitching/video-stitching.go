package problem1024

import "sort"

func videoStitching(clips [][]int, T int) int {
	sort.Slice(clips, func(i int, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})

	l, r := clips[0][0], clips[0][1]
	if l != 0 {
		return -1
	}

	res := 1
	for i := 1; i < len(clips); i++ {
		if r < clips[i][0] {
			return -1
		}
		if r < clips[i][1] {
			r = clips[i][1]
			res++
			if r >= T {
				break
			}
		}
	}

	if r < T {
		return -1
	}
	return res
}
