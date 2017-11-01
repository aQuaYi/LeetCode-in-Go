package Problem0390

func lastRemaining(n int) int {
	isLeftStart := true
	// remain 是剩余数字的个数
	remain := n

	// 剩余数字中，最小的数
	min := 1
	// min + step == 第二小的数字
	step := 1
	for remain > 1 {
		if isLeftStart || remain%2 == 1 {
			min += step
		}

		isLeftStart = !isLeftStart
		step *= 2
		remain /= 2
	}

	return min
}
