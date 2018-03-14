package problem0084

func largestRectangleArea(heights []int) int {
	// 在 h 结尾添加 -1 可以让 for 循环中，求解 area 的逻辑一致
	h := append(heights, -1)
	n := len(h)

	var maxArea, height, left, right, area int
	// 如果已知heights数组是升序的，应该怎么做？
	// 比如1,2,5,7,8
	// 那么就是(1*5) vs. (2*4) vs. (5*3) vs. (7*2) vs. (8*1)
	// 也就是max(height[i]*(size-i))
	// 使用栈的目的就是构造这样的升序序列，按照以上方法求解。
	var stack []int
	// stack 中存的是 h 的元素的索引号
	// stack 中索引号对应的 h 中的值，是递增的。
	// e.g.
	//     stack = []int{1,3,5}，那么
	//     h[1] <= h[3] <= h[5]

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

		// area = h[left+1:right] * min(h[left+1:right])
		area = (right - left - 1) * height
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
