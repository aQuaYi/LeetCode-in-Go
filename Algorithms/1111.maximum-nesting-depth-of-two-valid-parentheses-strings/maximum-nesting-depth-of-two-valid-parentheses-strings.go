package problem1111

func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	stack, top := make([]int, len(seq)), -1
	for i, r := range seq {
		if r == '(' {
			top++
			stack[top], res[i] = i, top%2
		} else {
			res[i] = res[stack[top]]
			top--
		}
	}
	return res
}
