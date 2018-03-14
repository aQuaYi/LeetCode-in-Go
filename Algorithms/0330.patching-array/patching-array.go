package problem0330

// nums 是升序排列的
func minPatches(nums []int, n int) (count int) {
	var max, i int

	// nums[:i] 的所有 子数组 的和的集合，等于 <max 内的所有自然数的集合
	max = 1

	for max <= n {
		if i < len(nums) && nums[i] <= max {
			// 根据 max 的定义
			// nums[:i]   --> max
			// nums[:i+1] --> max+nums[i]
			max += nums[i]
			i++
		} else {
			// 此时
			// nums[i] > max
			// 如果还 max += nums[i] 的话
			// 会导致 x ∈ [max, nums[i]) 无法获取
			// 需要添加 max 来实现，来让 自然数的集合达到 [1, 2*max)
			max <<= 1
			count++
		}
	}

	return
}
