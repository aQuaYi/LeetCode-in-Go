package problem0307

// NumArray 是 nums 的和的切片
type NumArray struct {
	nums []int
}

// Constructor 返回 NumArray
func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

// Update 更新 nums
func (na *NumArray) Update(i int, v int) {
	na.nums[i] = v
}

// SumRange 返回 sum(nums[i:j+1])
func (na *NumArray) SumRange(i int, j int) int {
	res := 0
	for k := i; k <= j; k++ {
		res += na.nums[k]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
