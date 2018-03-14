package problem0525

func findMaxLength(nums []int) int {
	delta := make(map[int]int, len(nums)/2)
	delta[0] = -1

	ones, zeros, res := 0, 0, 0

	for i := range nums {
		if nums[i] == 1 {
			ones++
		} else {
			zeros++
		}

		if j, ok := delta[ones-zeros]; ok {
			res = max(res, i-j)
		} else {
			delta[ones-zeros] = i
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
