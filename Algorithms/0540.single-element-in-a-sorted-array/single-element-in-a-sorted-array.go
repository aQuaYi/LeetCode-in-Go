package problem0540

func singleNonDuplicate(nums []int) int {
	n := len(nums) - 1
	lo, hi := 0, n-1

	// 二分查找
	for lo < hi {
		mid := lo + (hi-lo)/2
		// 为了简化后面的判断
		// 让 mid 总是落在偶数位上
		// 这样的话，在遇到单独的数之前，总有 nums[mid] == nums[mid+1]
		if mid%2 == 1 {
			mid--
		}

		if nums[mid] != nums[mid+1] {
			// 单独的数，在 mid 之前
			hi = mid
			// hi 没有 -1 是因为
			// 此时的 mid 可能就是答案，-1 的话，有可能会漏掉答案
		} else {
			// 单独的数，在 mid 之后
			lo = mid + 2
			// nums[:mid+2] 中肯定没有包含答案，所以，把 mid + 2 放入 lo
		}

	}

	return nums[lo]
}
