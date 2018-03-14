package problem0055

func canJump(nums []int) bool {

	for i := len(nums) - 2; i >= 0; i-- {
		// 找到数值为 0 的元素
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if i-j < nums[j] {
				// 在 j 号位置上，可以跨过 0 元素
				i = j
				break
			}
		}

		if j == -1 {
			// 在 0 元素之前，没有位置可以跨过 0
			return false
		}
	}

	return true
}
