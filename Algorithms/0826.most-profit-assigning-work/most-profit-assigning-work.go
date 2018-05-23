package problem0826

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs := make([]job, len(difficulty))
	for i := range difficulty {
		jobs[i] = job{
			d: difficulty[i],
			p: profit[i],
		}
	}
	sort.Slice(jobs, func(i int, j int) bool {
		return jobs[i].d < jobs[j].d
	})

	sort.Ints(worker)

	res := 0
	i := 0
	maxp := 0
	for _, ability := range worker {
		for i < len(jobs) && ability >= jobs[i].d {
			maxp = max(maxp, jobs[i].p)
			i++
		}
		res += maxp
	}

	return res
}

type job struct {
	d, p int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
