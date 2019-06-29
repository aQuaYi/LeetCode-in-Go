package problem1037

func isBoomerang(points [][]int) bool {
	return isAllDistinct(points) && !isLine(points)
}

func isAllDistinct(points [][]int) bool {
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if isEqual(points[i], points[j]) {
				return false
			}
		}
	}
	return true
}

func isEqual(p, q []int) bool {
	return p[0] == q[0] && p[1] == q[1]
}

func isLine(points [][]int) bool {
	a, b, c := points[0], points[1], points[2]
	h0, w0 := a[1]-b[1], a[0]-b[0]
	h1, w1 := a[1]-c[1], a[0]-c[0]
	return h0*w1 == h1*w0
}
