package Problem0643

func findMaxAverage(nums []int, k int) float64 {
	temp := 0
	for j := 0; j < k; j++ {
		temp += nums[j]
	}

	max := temp

	for i, j := 0, k; j < len(nums); i, j = i+1, j+1 {
		temp = temp - nums[i] + nums[j]

		if max < temp {
			max = temp
		}
	}

	return float64(max) / float64(k)
}
