package problem0056

import (
	"sort"
)

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(its []Interval) []Interval {
	if len(its) <= 1 {
		return its
	}

	sort.Slice(its, func(i int, j int) bool {
		return its[i].Start < its[j].Start
	})

	res := make([]Interval, 0, len(its))

	temp := its[0]
	for i := 1; i < len(its); i++ {
		if its[i].Start <= temp.End {
			temp.End = max(temp.End, its[i].End)
		} else {
			res = append(res, temp)
			temp = its[i]
		}
	}
	res = append(res, temp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
