package problem0836

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 把 rec 分别投影到水平轴和竖直轴上
	h1, v1 := projecting(rec1)
	h2, v2 := projecting(rec2)
	// 两个投影线同时相交可以认为 rec 相交
	return isLineOverlap(h1, h2) && isLineOverlap(v1, v2)
}

func projecting(rec []int) (h, v []int) {
	h = []int{rec[0], rec[2]}
	v = []int{rec[1], rec[3]}
	return
}

func isLineOverlap(a, b []int) bool {
	// 没有分开，就是相交
	return !(a[1] <= b[0] || b[1] <= a[0])
}
