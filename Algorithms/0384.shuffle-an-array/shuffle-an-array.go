package problem0384

import (
	"math/rand"
)

// Solution 是答案
type Solution struct {
	origNums, nums []int
}

// Constructor 构建 Solution
func Constructor(nums []int) Solution {
	o := make([]int, len(nums))
	n := make([]int, len(nums))
	copy(o, nums)
	copy(n, nums)
	return Solution{origNums: o, nums: n}
}

// Reset the array to its original configuration and return it.
func (s Solution) Reset() []int {
	return s.origNums
}

// Shuffle returns a random shuffling of the array.
func (s Solution) Shuffle() []int {
	i, j := len(s.nums), 0
	for 1 < i {
		j = rand.Intn(i)
		s.nums[i-1], s.nums[j] = s.nums[j], s.nums[i-1]
		i--
	}

	return s.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
