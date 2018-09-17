package problem0033

func search(nums []int, target int) int {
	size := len(nums)

	lo, hi := 0, size-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	/* lo = hi，是最小值的索引值 */

	rotated := lo /* 数组旋转了的距离 */

	lo, hi = 0, size-1

	for lo <= hi {
		mid := (lo + hi) / 2
		/* nums 是 rotated，所以需要使用 rotatedMid 来获取 mid 的值 */
		rotatedMid := (rotated + mid) % size
		switch {
		case nums[rotatedMid] > target:
			hi = mid - 1
		case nums[rotatedMid] < target:
			lo = mid + 1
		default:
			return rotatedMid
		}
	}

	return -1
}
