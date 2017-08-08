package Problem0033

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if nums[len(nums)-1] > nums[0] {
		return searchInt(nums, target, 0)
	}

	base := searchMax(nums) + 1
	res := searchInt(nums[:base], target, 0)
	if res != -1 {
		return res
	}
	res = searchInt(nums[base:], target, base)
	if res != -1 {
		return res
	}

	return -1
}

func searchInt(nums []int, target, base int) (index int) {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		index = (lo + hi) / 2
		switch {
		case nums[index] > target:
			hi = index - 1
		case nums[index] < target:
			lo = index + 1
		default:
			return base + index
		}
	}
	return -1
}

func searchMax(nums []int) (index int) {
	i, j := 0, len(nums)-1

	for {
		index = (i + j) / 2
		if index == 0 || index == len(nums)-1 {
			return
		}
		if index > 0 && index < len(nums)-1 && nums[index] > nums[index-1] && nums[index] > nums[index+1] {
			return
		}
		if nums[index] > nums[i] {
			i = index
		} else {
			j = index
		}
	}
}
