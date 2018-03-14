package problem0598

func maxCount(m int, n int, ops [][]int) int {
	for _, o := range ops {
		m = min(m, o[0])
		n = min(n, o[1])
	}
	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
