package Problem0238

func productExceptSelf(a []int) []int {
	l := len(a)
	// left[i] 是 a[i] 左侧所有元素的乘积
	// right[i] 是 a[i] 右侧所有元素的乘积
	left, right := make([]int, l), make([]int, l)

	// 题目已经保证了 n >= 2
	left[0], right[l-1] = 1, 1
	left[1], right[l-2] = a[0], a[l-1]

	for i := 2; i < l; i++ {
		left[i] = a[i-1] * left[i-1]
		right[l-i-1] = a[l-i] * right[l-i]
	}

	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}

	return res
}
