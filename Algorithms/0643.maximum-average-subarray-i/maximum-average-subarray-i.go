package problem0643

func findMaxAverage(nums []int, k int) float64 {
	temp := 0
	// 获取第一个temp
	for i := 0; i < k; i++ {
		temp += nums[i]
	}

	max := temp

	for i := k; i < len(nums); i++ {
		// 根据 i 获取 temp
		temp = temp - nums[i-k] + nums[i]

		if max < temp {
			max = temp
		}
	}

	return float64(max) / float64(k)
}
