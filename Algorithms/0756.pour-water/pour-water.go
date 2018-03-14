package problem0756

func pourWater(heights []int, V int, K int) []int {
	for V > 0 {
		V--
		if !isDroppedLeft(heights, K) && !isDroppedRight(heights, K) {
			heights[K]++
		}
	}
	return heights
}

// 当 water 可以放在 cent 的 left 边时，返回 true
// 否则，返回 false
func isDroppedLeft(heights []int, center int) bool {
	index := center
	i := center
	for 0 <= i-1 && heights[i-1] <= heights[i] {
		// 出现落差，说明
		// i-1 是一个可行的落脚点
		if heights[i-1] < heights[i] {
			index = i - 1
		}
		i--
	}
	if index < center {
		heights[index]++
		return true
	}
	return false
}

// 当 water 可以放在 cent 的 right 边时，返回 true
// 否则，返回 false
func isDroppedRight(heights []int, center int) bool {
	index := center
	i := center
	for i+1 < len(heights) && heights[i] >= heights[i+1] {
		// 出现落差，说明
		// i+1 是一个可行的落脚点
		if heights[i] > heights[i+1] {
			index = i + 1
		}
		i++
	}
	if center < index {
		heights[index]++
		return true
	}
	return false
}
