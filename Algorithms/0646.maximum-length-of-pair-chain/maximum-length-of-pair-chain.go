package Problem0646

import (
	"sort"
)

func findLongestChain(ps [][]int) int {
	sort.Sort(pairs(ps))
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

// pairs 实现了 sort.Interface 接口
type pairs [][]int

func (p pairs) Len() int { return len(p) }

func (p pairs) Less(i, j int) bool {
	if p[i][1] == p[j][1] {
		return p[i][0] > p[j][0]
	}
	return p[i][1] < p[j][1]
}

func (p pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
