package Problem0053

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i; j < len(nums); j++ {
			temp += nums[j]
			if temp > max {
				max = temp
			}
		}
	}

	return max
}
