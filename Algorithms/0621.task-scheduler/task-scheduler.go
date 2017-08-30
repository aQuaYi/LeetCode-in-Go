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
	remain := 0
	for len(rec) > 0 {
		if remain > 0 {
			res += remain
		}

		remain = n

		begin := len(rec) - 1
		for rec[begin] != 0 {
			rec[begin]--
			remain--
			begin--
		}

		rec = rec[begin+1:]
	}

	return res
}
