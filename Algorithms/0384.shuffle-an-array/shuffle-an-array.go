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
	return Solution{
		origNums: nums,
		nums:     append([]int{}, nums...),
	}
}

// Reset the array to its original configuration and return it.
func (this *Solution) Reset() []int {
	return this.origNums
}

// Shuffle returns a random shuffling of the array.
func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.nums); i++ {
		r := rand.Intn(len(this.nums))
		this.nums[r], this.nums[i] = this.nums[i], this.nums[r]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
