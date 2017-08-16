package Problem0643

func findMaxAverage(nums []int, k int) float64 {
	temp := 0
	for i := 0; i < k; i++ {
		temp += nums[i]
	}

	max := temp

	for i := k; i < len(nums); i++ {
		temp = temp - nums[i-k] + nums[i]

		if max < temp {
			max = temp
		}
	}

	return float64(max) / float64(k)
}
