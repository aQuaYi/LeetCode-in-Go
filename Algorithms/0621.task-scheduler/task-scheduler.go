package Problem0621

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	rec := make([]int, 26)
	for _, b := range tasks {
		rec[b-'A']++
	}
	sort.Ints(rec)

	res := 0
	intervals := 0
	idx := 0
	for len(rec) > 0 {
		if intervals > 0 {
			res += intervals
		}

		// 新周期
		intervals = n + 1

		idx = len(rec) - 1
		for intervals > 0 && idx >= 0 && rec[idx] > 0 {
			rec[idx]--
			intervals--
			idx--
			res++
		}

		sort.Ints(rec)
		idx = len(rec) - 1
		for idx >= 0 && rec[idx] > 0 {
			idx--
		}

		rec = rec[idx+1:]
	}

	return res
}
