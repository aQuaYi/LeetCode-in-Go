package Problem0011

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			temp := h(height, i, j) * w(i, j)
			if temp > max {
				max = temp
			}
		}
	}

	return max
}

// h 高度
func h(h []int, i, j int) int {
	a, b := h[i], h[j]

	if a < b {
		return a
	}

	return b
}

// w 宽度
func w(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}
