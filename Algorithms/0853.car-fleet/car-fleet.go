package problem0853

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	size := len(position)
	rs := make([]record, size)
	for i := range rs {
		p, s := position[i], speed[i]
		rs[i] = record{
			initPos:     p,
			arrivalTime: float64(target-p) / float64(s),
		}
	}

	// initPos 越大的车，越早到
	sort.Slice(rs, func(i int, j int) bool {
		return rs[i].initPos > rs[j].initPos
	})

	fleetTime := 0.
	res := 0

	for _, r := range rs {
		at := r.arrivalTime
		if fleetTime < at {
			fleetTime = at
			res++
		}
	}

	return res
}

type record struct {
	initPos     int
	arrivalTime float64
}
