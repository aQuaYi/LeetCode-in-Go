package problem0939

const initialArea = 2000000000 // > 40000*40000

func minAreaRect(points [][]int) int {
	size := len(points)

	isExisting := make(map[[2]int]bool, size)
	for _, p := range points {
		x, y := p[0], p[1]
		isExisting[[2]int{x, y}] = true
	}

	minArea := initialArea

	for i := 0; i < size; i++ {
		xi, yi := points[i][0], points[i][1]
		for j := i + 1; j < size; j++ {
			xj, yj := points[j][0], points[j][1]
			if xi == xj || yi == yj {
				continue
			}
			area := abs((xi - xj) * (yi - yj))
			if area >= minArea || // NOTICE: delay heavy operation
				!isExisting[[2]int{xi, yj}] ||
				!isExisting[[2]int{xj, yi}] {
				continue
			}
			minArea = area
		}
	}

	if minArea == initialArea {
		return 0
	}
	return minArea
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
