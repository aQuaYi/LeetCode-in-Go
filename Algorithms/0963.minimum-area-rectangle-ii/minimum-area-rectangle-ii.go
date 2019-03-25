package problem0963

import "math"

func minAreaFreeRect(points [][]int) float64 {
	size := len(points)

	// collect point couple which has the same middle point
	couples := make(map[[2]int][][2]int, size)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			// x, y := float64(xi+xj)/2, float64(yi+yj)/2
			// change for speedup
			x, y := xi+xj, yi+yj
			m := [2]int{x, y}
			couples[m] = append(couples[m], [2]int{i, j})
		}
	}

	res := math.MaxFloat64
	for _, c := range couples {
		size := len(c)
		for i := 0; i < size; i++ {
			p, q := points[c[i][0]], points[c[i][1]]
			for j := i + 1; j < size; j++ {
				o := points[c[j][0]]
				res = min(res, area(p, q, o))
			}
		}
	}

	if res == math.MaxFloat64 {
		return 0
	}
	return res
}

func area(p, q, o []int) float64 {
	xop, yop := p[0]-o[0], p[1]-o[1] // vector form o to p
	xoq, yoq := q[0]-o[0], q[1]-o[1] // vector form o to q
	if xop*xoq+yop*yoq != 0 {        // not a rectangle
		return math.MaxFloat64
	}
	return length(xop, yop) * length(xoq, yoq)
}

// length of vector (x,y)
func length(x, y int) float64 {
	xf, yf := float64(x), float64(y)
	return math.Sqrt(xf*xf + yf*yf)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
