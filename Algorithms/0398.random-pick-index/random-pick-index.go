package problem0398

import "math/rand"

// Solution 是答案所需的数据结构
type Solution struct {
	m map[int][]int
}

// Constructor 构建了 Solution
func Constructor(nums []int) Solution {
	m := make(map[int][]int, len(nums))
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	return Solution{m: m}
}

// Pick 随机出现 target 在 nums 中的索引号
func (s *Solution) Pick(target int) int {
	return s.m[target][rand.Intn(len(s.m[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
