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

func (s Solution) Reset() []int {
	return s.origNums
}

func (s Solution) Shuffle() []int {
	i, j := len(s.origNums), 0
	for 1 < i {
		j = rand.Intn(i)
		s.nums[i-1], s.nums[j] = s.nums[j], s.nums[i-1]
		i--
	}
	return s.nums
}

