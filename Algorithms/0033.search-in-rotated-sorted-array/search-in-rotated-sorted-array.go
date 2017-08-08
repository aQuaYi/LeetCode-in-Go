package Problem0033

func search(nums []int, target int) int {
	var index, indexOfMax int

	length := len(nums)

	if length == 0 {
		return -1
	}

	for indexOfMax+1 < length && nums[indexOfMax] < nums[indexOfMax+1] {
		indexOfMax++
	}

	lo, hi, median := 0, length-1, 0
	for lo <= hi {
		median = (lo + hi) / 2

		relativeIndex := median + indexOfMax + 1
		if relativeIndex < length {
			index = relativeIndex
		} else {
			index = relativeIndex - length
		}

		switch {
		case nums[index] > target:
			hi = median - 1
		case nums[index] < target:
			lo = median + 1
		default:
			return index
		}
	}

	return -1
}
