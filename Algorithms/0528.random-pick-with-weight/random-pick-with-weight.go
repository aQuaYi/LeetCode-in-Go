package problem0528

import (
	"math/rand"
	"sort"
)

// Solution object will be instantiated and called as such:
// obj := Constructor(w);
// param_1 := obj.PickIndex();
type Solution struct {
	counts []int
	total  int
}

// Constructor 返回 Solution
func Constructor(w []int) Solution {
	total := 0
	counts := make([]int, len(w))
	for i := range w {
		total += w[i]
		counts[i] = total
	}

	return Solution{
		counts: counts,
		total:  total,
	}
}

// PickIndex 根据权重返回 索引号
func (s *Solution) PickIndex() int {
	cand := rand.Intn(s.total)
	index := sort.Search(len(s.counts), func(i int) bool {
		return s.counts[i] > cand
	})
	return index
}
