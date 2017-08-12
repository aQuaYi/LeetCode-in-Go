package Problem0042

func trap(height []int) int {
	height = trim(height)
	res := 0

	for len(height) > 2 {
		for i := 1; i < len(height)-1; i++ {
			if hasHisherBefore(height, i) && hasHisherAfter(height, i) {
				height[i]++
				res++
			}
		}
		height = down(height)
	}

	return res
}

func hasHisherBefore(nums []int, index int) bool {
	for i := index - 1; i >= 0; i-- {
		if nums[i] > nums[index] {
			return true
		}
	}
	return false
}
func hasHisherAfter(nums []int, index int) bool {
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] {
			return true
		}
	}
	return false
}

func trim(nums []int) []int {
	i := 0
	length := len(nums)
	for i < len(nums) {
		if nums[i] != 0 {
			nums = nums[i:]
			break
		}
		i++
	}

	if i == length {
		return []int{}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			nums = nums[:i+1]
			break
		}
	}
	return nums
}

func down(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[i]--
		}
	}

	return trim(nums)
}
