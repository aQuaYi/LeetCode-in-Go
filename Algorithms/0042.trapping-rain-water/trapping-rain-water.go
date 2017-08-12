package Problem0042

func trap(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		wall := smaller(left(height, i), right(height, i))
		if wall > height[i] {
			res += wall - height[i]
		}
	}

	return res
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func left(nums []int, index int) int {
	res := nums[index]
	for i := index - 1; i >= 0; i-- {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func right(nums []int, index int) int {
	res := nums[index]
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
