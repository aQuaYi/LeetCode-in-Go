package problem0223

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	return area(A, B, C, D) + area(E, F, G, H) - area(max(A, E), max(B, F), min(C, G), min(D, H))
}

func area(A, B, C, D int) int {
	return edge(A, C) * edge(B, D)
}

func edge(A, C int) int {
	// 边长应该大于 0
	// 这样做就统一了两个正方形相交或者不相交的情况
	return max(0, C-A)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
