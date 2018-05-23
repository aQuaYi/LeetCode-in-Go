package problem0826

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	works := make([]work, len(difficulty))
	for i := range difficulty {
		works[i] = work{
			d: difficulty[i],
			p: profit[i],
		}
	}

	sort.Slice(works, func(i int, j int) bool {
		if works[i].p == works[j].p {
			return works[i].d < works[j].d
		}
		return works[i].p > works[j].p
	})

	res := 0
	for _, w := range worker {
		for i := range works {
			if works[i].d <= w {
				res += works[i].p
				break
			}
		}
	}
	return res
}

type work struct {
	d, p int
}
