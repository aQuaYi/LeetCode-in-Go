package problem0646

import (
	"sort"
)

func findLongestChain(ps [][]int) int {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i][1] == ps[j][1] {
			return ps[i][0] > ps[j][0]
		}
		return ps[i][1] < ps[j][1]
	})

	res := 0
	b := -1 << 32
	for i := 0; i < len(ps); i++ {
		c := ps[i][0]
		if b < c {
			res++
			b = ps[i][1]
		}
	}

	return res
}
