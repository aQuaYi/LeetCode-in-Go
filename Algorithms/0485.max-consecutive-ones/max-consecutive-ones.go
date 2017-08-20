package Problem0485

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			temp = 0
		}
		if nums[i] == 1 {
			temp++
			if max < temp {
				max = temp
			}
		}
	}
	return max
}
