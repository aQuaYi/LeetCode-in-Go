package Problem0756

func pourWater(heights []int, V int, K int) []int {
	for V > 0 {
		V--
		if !isDroppedLeft(heights, K) && !isDroppedRight(heights, K) {
			heights[K]++
		}
	}
	return heights
}

func isDroppedLeft(heights []int, center int) bool {
	idx := center
	i := center
	for 0 <= i-1 && heights[i-1] <= heights[i] {
		if heights[i-1] < heights[i] {
			idx = i - 1
		}
		i--
	}
	if idx < center {
		heights[idx]++
		return true
	}
	return false
}

func isDroppedRight(heights []int, center int) bool {
	idx := center
	i := center
	for i+1 < len(heights) && heights[i] >= heights[i+1] {
		if heights[i] > heights[i+1] {
			idx = i + 1
		}
		i++
	}
	if center < idx {
		heights[idx]++
		return true
	}
	return false
}
