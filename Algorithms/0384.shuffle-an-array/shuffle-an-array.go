package Problem0384

// Solution 是答案
type Solution struct {
	original []int
	rec      map[int]bool
}

// Constructor 构建 Solution
func Constructor(nums []int) Solution {
	o := make([]int, len(nums))
	copy(o, nums)
	r := make(map[int]bool, len(nums))
	for _, n := range nums {
		r[n] = true
	}
	return Solution{original: o, rec: r}
}

// Reset the array to its original configuration and return it.
func (s Solution) Reset() []int {
	return s.original
}

// Shuffle returns a random shuffling of the array.
func (s Solution) Shuffle() []int {
	temp := make([]int, 0, len(s.original))
	for n := range s.rec {
		temp = append(temp, n)
	}
	return temp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
