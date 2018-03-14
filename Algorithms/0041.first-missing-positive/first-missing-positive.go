package problem0041

func firstMissingPositive(nums []int) int {
	// 整理 nums ，让 nums[k] == k+1，只要 k+1 存在于 nums 中
	for i := 0; i < len(nums); i++ {

		// fmt.Println(i)

		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			// 当 for 的判断语句成立时，
			// nums[i]-1 就是 k ，nums[i] 的值是 k+1
			// nums[i] != nums[nums[i]-1] 即是 k+1 != nums[k] ，这说明
			// 1. k+1 存在与 nums 中，
			// 2. k+1 还没有在他该在的 nums[k] 中
			// 通过互换，让 k+1 到 nums[k] 中去
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

			// fmt.Println(nums)

			// 使用 for 而不是 if 是因为
			// nums[i] 中的新值，有可能是另一个 k+1 ，需要再让其归为
			// 如果使用 if ，而这个新的 k+1 又只有一个的话，
			// 这个新的 k+1 不会被处理到，不会被放在 nums[k] 中
		}
	}
	// 循环结束后，所有 1<=k+1<=len(nums) 且 k+1 存在于nums中，都会被存放于 nums[k] 中

	// 整理后，第一个不存在的 k+1 就是答案
	for k := range nums {
		if nums[k] != k+1 {
			return k + 1
		}
	}
	return len(nums) + 1
}
