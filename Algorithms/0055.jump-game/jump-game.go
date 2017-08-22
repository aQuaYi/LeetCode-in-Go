package Problem0055

func canJump(nums []int) bool {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if i-j < nums[j] {
				i = j
				break
			}
		}

		if j == -1 {
			return false
		}
	}

	return true
}
