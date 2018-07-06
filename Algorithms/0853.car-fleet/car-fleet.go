package problem0853

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	size := len(position)
	rs := make([]record, size)
	for i := range rs {
		pos, spd := position[i], speed[i]
		rs[i] = record{
			initPos:     pos,
			arrivalTime: float64(target-pos) / float64(spd),
		}
	}

	sort.Slice(rs, func(i int, j int) bool {
		return rs[i].initPos > rs[j].initPos
	})

	fleetTime := -1.0
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
