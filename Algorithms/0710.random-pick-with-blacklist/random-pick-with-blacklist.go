package problem0710

import (
	"math/rand"
	"time"
)

// Solution 包含了 BlackList 和 N
type Solution struct {
	M        int
	blackMap map[int]int
}

// Constructor 构建了 Solution
func Constructor(N int, blacklist []int) Solution {
	rand.Seed(time.Now().UnixNano())

	M := N - len(blacklist)

	blackMap := make(map[int]int, len(blacklist))
	//
	for _, b := range blacklist {
		blackMap[b] = -1
	}
	for _, b := range blacklist {
		if b >= M {
			continue
		}

		for blackMap[N-1] == -1 {
			N--
		}

		blackMap[b] = N - 1
		N--
	}

	return Solution{
		M:        M,
		blackMap: blackMap,
	}

}

// Pick 选取了不在 BlackList 中的值
func (s *Solution) Pick() int {
	r := rand.Intn(s.M)
	if t, ok := s.blackMap[r]; ok {
		return t
	}
	return r
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
