package problem0939

const initialArea = 2000000000 // > 40000*40000

func minAreaRect(points [][]int) int {
	size := len(points)

	isExisting := make(map[[2]int]bool, size)
	for _, p := range points {
		x, y := p[0], p[1]
		isExisting[[2]int{x, y}] = true
	}

	area := initialArea

	for i := 0; i < size; i++ {
		x0, y0 := points[i][0], points[i][1]
		for j := i + 1; j < size; j++ {
			x1, y1 := points[j][0], points[j][1]
			if x0 == x1 || y0 == y1 {
				continue
			}
			newArea := abs(x0-x1) * abs(y0-y1)
			if newArea > area || // NOTICE: delay heavy operation
				!isExisting[[2]int{x1, y0}] ||
				!isExisting[[2]int{x0, y1}] {
				continue
			}
			area = newArea
		}
	}

	if area == initialArea {
		return 0
	}
	return area
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
