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
	// M 是可以自由选择的数的个数
	M := N - len(blacklist)

	blackMap := make(map[int]int, len(blacklist))
	// 由于 blacklist 是乱序的，只好先把 blacklist 中的元素全部添加为 blackMap 的 key
	for _, b := range blacklist {
		blackMap[b] = b
	}

	for _, b := range blacklist {
		if b >= M {
			continue
		}

		// 对于所有的 b < M
		// 与 [M,N) 中某个不在 BlackList 中的数，关联上

		N--
		for blackMap[N] == N {
			N--
		}

		blackMap[b] = N
	}

	return Solution{
		M:        M,
		blackMap: blackMap,
	}

}

// Pick 选取了不在 BlackList 中的值
func (s *Solution) Pick() int {
	// 在 [0,M) 中任意挑选一个数 r
	r := rand.Intn(s.M)
	if t, ok := s.blackMap[r]; ok {
		// r 是 s.blackMap 的 key
		// 说明 r 是 BlackList 中的值
		// 需要返回 与 r 相关联的值
		return t
	}
	return r
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
