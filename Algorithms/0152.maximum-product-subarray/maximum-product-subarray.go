package Problem0152

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	size := len(nums)

	maxp_sub_arr := make([]int, size)
	ng_maxp_sub_arr := make([]int, size)
	maxp_sub_arr[0] = nums[0]
	ng_maxp_sub_arr[0] = 1
	if maxp_sub_arr[0] < 0 {
		ng_maxp_sub_arr[0] = maxp_sub_arr[0]
	}

	for i := 1; i < size; i += 1 {
		if nums[i] == 0 {
			continue
		}

		if maxp_sub_arr[i-1] == 0 {
			maxp_sub_arr[i] = nums[i]
			if nums[i] < 0 {
				ng_maxp_sub_arr[i] = maxp_sub_arr[i]
			} else {
				ng_maxp_sub_arr[i] = 1
			}
		}

		if nums[i] > 0 {
			maxp_sub_arr[i] = max(nums[i]*maxp_sub_arr[i-1], nums[i])
			ng_maxp_sub_arr[i] = nums[i] * ng_maxp_sub_arr[i-1]
		}

		if nums[i] < 0 {
			maxp_sub_arr[i] = nums[i] * ng_maxp_sub_arr[i-1]
			ng_maxp_sub_arr[i] = min(nums[i]*maxp_sub_arr[i-1], nums[i])
		}
	}

	maxp := maxp_sub_arr[0]
	for _, subp := range maxp_sub_arr {
		if subp > maxp {
			maxp = subp
		}
	}

	return maxp
}
