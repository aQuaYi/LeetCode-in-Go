package Problem0307

// NumArray 是 nums 的和的切片
type NumArray struct {
	nums, dp []int
}

// Constructor 返回 NumArray
func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i+1] = dp[i] + nums[i]
	}

	return NumArray{nums: nums, dp: dp}
}

// Update 更新 nums
func (na *NumArray) Update(i int, v int) {
	for idx := i + 1; idx < len(na.dp); idx++ {
		na.dp[idx] += v - na.nums[i]
	}
}

// SumRange 返回 sum(nums[i:j+1])
func (na *NumArray) SumRange(i int, j int) int {
	return na.dp[j+1] - na.dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
