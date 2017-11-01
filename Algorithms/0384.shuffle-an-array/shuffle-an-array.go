package Problem0384

import (
	"math/rand"
)

// Solution 是答案
type Solution struct {
	original []int
}

// Constructor 构建 Solution
func Constructor(nums []int) Solution {
	o := make([]int, len(nums))
	copy(o, nums)
	return Solution{original: o}
}

// Reset the array to its original configuration and return it.
func (s Solution) Reset() []int {
	return s.original
}

// Shuffle returns a random shuffling of the array.
func (s Solution) Shuffle() []int {
	temp := make([]int, len(s.original))
	copy(temp, s.original)

	i, j := len(s.original), 0
	for 1 < i {
		j = rand.Intn(i)
		temp[i-1], temp[j] = temp[j], temp[i-1]
		i--
	}

	return temp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
