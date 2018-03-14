package problem0303

// NumArray 是切片的和切片
type NumArray struct {
	dp []int
}

// Constructor 返回 NumArray
func Constructor(nums []int) NumArray {
	size := len(nums)
	// dp[i+1] == SumRange(0,i)
	dp := make([]int, size+1)
	for i := 1; i <= size; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}

	return NumArray{dp: dp}
}

// SumRange (i,j) == sum(nums[i:j+1])
func (na *NumArray) SumRange(i int, j int) int {
	return na.dp[j+1] - na.dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
