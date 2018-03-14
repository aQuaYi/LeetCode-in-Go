package problem0169

func majorityElement(nums []int) int {
	// 根据题意 len[nums] > 0 且 出现次数大于 n/2 的元素存在。
	x, t := nums[0], 1

	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			// 此时 x != nums[i] 且 t == 0
			// 可知 i 必定为 偶数
			// 假设 nums 中出现最多的元素是 z，其出现次数为 zn > n/2
			// 在 nums[:i] 中，z 出现次数 <=i/2
			// 那么，在 nums[i:] 中，z 出现的次数 >= zn - i/2 > n/2 - i/2 = (n-i)/2
			// 即 z 在 nums[i:] 中出现的次数，依然超过了 len(nums[i:])/2
			x = nums[i]
			t = 1
		}
	}

	return x
}
