package problem0081

func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}

	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return true
		}
	}

	return false
}
