package problem0033

func search(nums []int, target int) int {
	rotated := indexOfMin(nums) /* 数组旋转了的距离 */
	size := len(nums)
	left, right := 0, size-1

	for left <= right {
		mid := (left + right) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid - 1
		default:
			return rotatedMid
		}
	}

	return -1
}

/* nums 是被旋转了的递增数组 */
func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := (left + right) / 2
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
