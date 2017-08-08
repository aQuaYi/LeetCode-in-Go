package Problem0033

func search(nums []int, target int) int {
	var index, indexOfMax int
	length := len(nums)

	if length == 0 {
		return -1
	}

	// 获取最大值的索引号，以便进行索引号变换
	for indexOfMax+1 < length && nums[indexOfMax] < nums[indexOfMax+1] {
		indexOfMax++
	}

	low, high, median := 0, length-1, 0
	for low <= high {
		median = (low + high) / 2

		// 变换索引号
		index = median + indexOfMax + 1
		if index >= length {
			index -= length
		}

		// 传统二分查找法的比较判断
		switch {
		case nums[index] > target:
			high = median - 1
		case nums[index] < target:
			low = median + 1
		default:
			return index
		}
	}

	return -1
}
