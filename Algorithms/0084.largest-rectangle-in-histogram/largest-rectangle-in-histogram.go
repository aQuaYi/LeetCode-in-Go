package Problem0084

func largestRectangleArea(h []int) int {
	h = append(h, -1)
	n := len(h)

	var maxArea, height, left, right, area int
	var stack []int

	for right < n {
		if len(stack) == 0 || h[stack[len(stack)-1]] <= h[right] {
			stack = append(stack, right)
			right++
			continue
		}

		height = h[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}

		area = (right - left - 1) * height
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
