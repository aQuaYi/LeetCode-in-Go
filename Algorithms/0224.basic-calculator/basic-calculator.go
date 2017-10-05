package Problem0224

// s 中不会含有 × ÷，所以，可以不用理会括号
func calculate(s string) int {
	nums := make([]int, 1, len(s))
	nums[0] = 0
	opt := '+'

	push := func(n int) {
		nums = append(nums, n)
	}

	pop := func() int {
		res := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return res
	}

	for _, b := range s {
		switch b {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			push(operate(pop(), int(b-'0'), opt))
		case '+', '-':
			opt = b
		}
	}

	return nums[0]
}

func operate(a, b int, opt rune) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		return 0
	}
}
