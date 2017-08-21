package Problem0283

func moveZeroes(nums []int) {
	indexZero := 0
	for indexZero < len(nums) && nums[indexZero] != 0 {
		indexZero++
	}

	i := indexZero + 1
	for i < len(nums) {
		if nums[i] != 0 {
			nums[i], nums[indexZero] = nums[indexZero], nums[i]
			for indexZero < len(nums) && nums[indexZero] != 0 {
				indexZero++
			}
		}
		i++
	}
}
