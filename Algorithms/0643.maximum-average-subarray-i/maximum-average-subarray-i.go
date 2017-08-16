package Problem0643

func findMaxAverage(nums []int, k int) float64 {
	max := 0

	for i := 0; i+k-1 < len(nums); i++ {
		temp := 0
		for j := i; j < i+k; j++ {
			temp += nums[j]
		}
		if max < temp {
			max = temp
		}
	}

	return float64(max) / float64(k)
}
