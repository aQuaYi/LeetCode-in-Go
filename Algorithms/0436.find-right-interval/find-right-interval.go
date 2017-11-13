package Problem0436

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type Interval = kit.Interval

func findRightInterval(intervals []Interval) []int {
	size := len(intervals)
	res := make([]int, size)

	for i := 0; i < size; i++ {
		e := intervals[i].End
		r := 1<<31 - 1
		res[i] = -1
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			if e <= intervals[j].Start && intervals[j].Start < r {
				res[i] = j
				r = intervals[j].Start
			}

			if e ==r {
				break
			}
		}
	}

	return res
}
