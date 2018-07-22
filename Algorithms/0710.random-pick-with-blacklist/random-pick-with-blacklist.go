package problem0710

import (
	"math/rand"
	"sort"
)

// Solution 包含了 BlackList 和 N
type Solution struct {
	N         int
	BlackList []int
}

// Constructor 构建了 Solution
func Constructor(N int, blacklist []int) Solution {
	sort.Ints(blacklist)
	return Solution{
		N:         N,
		BlackList: blacklist,
	}
}

// Pick 选取了不在 BlackList 中的值
func (s *Solution) Pick() int {
	for {
		res := rand.Intn(s.N)
		if !isInBlackList(s.BlackList, res) {
			return res
		}
	}
}

// bl 是升序排列
func isInBlackList(a []int, n int) bool {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		switch {
		case a[m] < n:
			l = m + 1
		case n < a[m]:
			r = m + 1
		default:
			return true
		}
	}
	return false
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
