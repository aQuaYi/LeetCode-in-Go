package problem0485

func findMaxConsecutiveOnes(nums []int) int {
	max := 0

	for i, j := 0, -1; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
		} else {
			if max < i-j {
				max = i - j
			}
		}
	}

	return max
}
