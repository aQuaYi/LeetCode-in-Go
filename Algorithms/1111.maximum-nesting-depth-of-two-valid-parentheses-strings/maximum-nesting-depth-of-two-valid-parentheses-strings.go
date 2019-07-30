package problem1111

func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	stack, top := make([]int, len(seq)), -1
	for i, r := range seq {
		if r == ')' {
			res[i] = res[stack[top]]
			top--
			continue
		}
		top++
		stack[top] = i
		res[i] = top % 2
	}

	return res
}
