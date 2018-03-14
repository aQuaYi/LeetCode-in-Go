package problem0391

type point struct {
	x, y int
}

func isRectangleCover(rects [][]int) bool {
	if len(rects) == 0 || len(rects[0]) == 0 {
		return false
	}

	// 记录最后合成的正方形的 四个顶点 的坐标
	minX, maxX := 1<<63-1, -1<<63
	minY, maxY := minX, maxX
	// 最后合成正方形的面积
	area := 0

	// 记录添加的正方形的 四个顶点 位置
	isCorner := make(map[point]bool)

	var p, p1, p2, p3, p4 point
	var x, y, x1, x2, y1, y2 int

	for _, r := range rects {
		x1, y1, x2, y2 = r[0], r[1], r[2], r[3]

		minX = min(x1, minX)
		minY = min(y1, minY)
		maxX = max(x2, maxX)
		maxY = max(y2, maxY)

		area += (x2 - x1) * (y2 - y1)

		for _, x = range []int{x1, x2} {
			for _, y = range []int{y1, y2} {
				p = point{x, y}

				// 如果想要最后能够合成一个矩形
				// 除了大矩形的四个顶点外，其余的小矩形的四个顶点会出现两次
				if isCorner[p] {
					delete(isCorner, p)
				} else {
					isCorner[p] = true
				}

			}
		}

	}

	p1 = point{minX, minY}
	p2 = point{minX, maxY}
	p3 = point{maxX, minY}
	p4 = point{maxX, maxY}

	if !isCorner[p1] ||
		!isCorner[p2] ||
		!isCorner[p3] ||
		!isCorner[p4] ||
		len(isCorner) != 4 {
		return false
	}

	return area == (maxX-minX)*(maxY-minY)
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
