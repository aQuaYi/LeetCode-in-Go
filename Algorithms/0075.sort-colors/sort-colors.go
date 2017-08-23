package Problem0075

func sortColors(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	if only2 := setHead(nums); only2 {
		return
	}

	// 三路快排
	i, j, k := 0, 1, length-1
	for j <= k {
		switch {
		case nums[j] < nums[i]:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case nums[i] < nums[j]:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		default:
			j++
		}
	}
	return
}

// 如果 nums 中有 1 ， nums[0] 应该为 1
// 否则，如果 nums 中还有 0， nums[0] 应该为 0
// 否则，说明 nums 中只有 2。
func setHead(nums []int) bool {
	i, l := 0, len(nums)
	// 设置 nums[0] = 1
	for i < l && nums[i] != 1 {
		i++
	}
	if i < l {
		nums[0], nums[i] = nums[i], nums[0]
		return false
	}

	// nums 中没有 1
	// 只好设置 nums[0] = 0
	i = 0
	for i < l && nums[i] != 0 {
		i++
	}

	if i < l {
		nums[0], nums[i] = nums[i], nums[0]
		return false
	}

	// nums 中只有 2
	return true
}
