package problem0645

func findErrorNums(nums []int) []int {
	// n == len(nums)
	// 而 nums 的数字属于 [1,n]，那么
	// 对于每一个 nums[abs(nums[i])-1] 中的数取反
	// 操作过程中，
	// 取反前，nums[abs(nums[i])-1] 已经是负数的话，说明 abs(nums[i]) 是重复的数
	// 操作结束后，
	// nums[i] > 0，说明 nums[i] 没有被取反过， i+1 是丢失的数
	dup := 0
	for i := 0; i < len(nums); i++ {
		n := abs(nums[i])

		if nums[n-1] < 0 {
			dup = n
			// NOTICE: nums[n-1]<0 的时候，不能再取反为正数了
		} else {
			nums[n-1] = -nums[n-1]
		}
	}

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
