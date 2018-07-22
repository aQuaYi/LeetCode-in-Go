package problem0710

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Solution 包含了 BlackList 和 N
type Solution struct {
	N         int
	BlackList []int
}

// Constructor 构建了 Solution
func Constructor(N int, blacklist []int) Solution {
	rand.Seed(time.Now().UnixNano())
	sort.Ints(blacklist)
	return Solution{
		N:         N,
		BlackList: blacklist,
	}
}

// Pick 选取了不在 BlackList 中的值
func (s *Solution) Pick() int {
	i := 0
	for {
		i++
		res := rand.Intn(s.N)
		if !isInBlackList(s.BlackList, res) {
			fmt.Printf("运行了 %d 次\n", i)
			return res
		}
	}
}

// bl 是升序排列
func isInBlackList(a []int, n int) bool {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		switch {
		case a[m] < n:
			l = m + 1
		case n < a[m]:
			r = m - 1
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
