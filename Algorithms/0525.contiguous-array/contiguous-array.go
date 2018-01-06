package Problem0525

func findMaxLength(nums []int) int {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	sumIndex := make(map[int]int, len(nums)/2)
	sumIndex[0] = -1

	sum, res := 0, 0

	for i := range nums {
		sum += nums[i]
		if index, ok := sumIndex[sum]; ok {
			res = max(res, i-index)
		} else {
			sumIndex[sum] = i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
