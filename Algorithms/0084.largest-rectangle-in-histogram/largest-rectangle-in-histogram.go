package problem0084

func largestRectangleArea(heights []int) int {
	// 在 heights 中原本是非负的数字
	// 再首尾添加了 -2 和 -1 后，简化 for 循环中 begin 的计算
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)

	// stack 中存的是 heights 的元素的索引号
	// stack 中索引号对应的 heights 中的值，是递增的。
	// e.g.
	//     stack = []int{1,3,5}，那么
	//     heights[1] < heights[3] < heights[5]
	stack := make([]int, 1, size)
	// 把 heights[0] 的索引号，已经放入了 stack
	// end 从 1 开始
	end := 1

	res := 0
	for end < size {
		// end 指向了新高，就把 end 放入 stack 后，指向下一个
		if heights[stack[len(stack)-1]] < heights[end] {
			stack = append(stack, end)
			end++
			continue
		}

		begin := stack[len(stack)-2]
		index := stack[len(stack)-1]
		height := heights[index]
		// area 是 heights(begin:end) 之间的最大方块的面积，因为
		// anyone of heights(begin:index) > height > anyone of heights(index:end)
		area := (end - begin - 1) * height

		res = max(res, area)

		// h 的索引号已经没有必要留在 stack 中了
		stack = stack[:len(stack)-1]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
