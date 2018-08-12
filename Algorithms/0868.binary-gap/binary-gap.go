package problem0868

func binaryGap(N int) int {
	res := 0
	gap := 0

	for N > 0 {
		if N&1 == 1 {
			res = max(res, gap)
			gap = 1
		} else if gap > 0 {
			// gap > 0 才 gap++ 是想要
			// 在遇到第一个 1 后，每次遇到 0 才 gap++
			gap++
		}
		N >>= 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
