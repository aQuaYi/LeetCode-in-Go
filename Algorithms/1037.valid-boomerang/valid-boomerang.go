package problem1037

func isBoomerang(p [][]int) bool {
	x0, y0 := p[0][0], p[0][1]
	x1, y1 := p[1][0], p[1][1]
	x2, y2 := p[2][0], p[2][1]
	return (x0-x2)*(y1-y2) != (y0-y2)*(x1-x2)
}
