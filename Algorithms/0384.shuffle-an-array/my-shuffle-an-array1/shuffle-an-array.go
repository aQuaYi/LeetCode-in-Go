package problem0384

import (
	"math/rand"
)


type Solution struct {
	origNums []int
	nums []int
}

func Constructor(nums []int) Solution {
	o := make([]int, len(nums))
	n := make([]int, len(nums))
	copy(o, nums)
	copy(n, nums)
	return Solution{origNums: o, nums: n}
}

func (this *Solution) Reset() []int {
	return this.origNums
}

func (this *Solution) Shuffle() []int {
	i, j := len(this.origNums), 0
	for 1 < i {
		j = rand.Intn(i)
		this.nums[i-1], this.nums[j] = this.nums[j], this.nums[i-1]
		i--
	}
	return this.nums
}

