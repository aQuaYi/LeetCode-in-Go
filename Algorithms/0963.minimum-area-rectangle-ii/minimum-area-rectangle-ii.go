package problem0963

import "math"

func minAreaFreeRect(points [][]int) float64 {
	mids := make(map[[2]int][][2]int, 0)
	size := len(points)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			x, y := xi+xj, yi+yj
			m := [2]int{x, y}
			mids[m] = append(mids[m], [2]int{i, j})
		}
	}

	res := math.MaxFloat64
	for _, ms := range mids {
		sz := len(ms)
		if sz == 1 {
			continue
		}
		for i := 0; i < sz; i++ {
			p, q := points[ms[i][0]], points[ms[i][1]]
			for j := i + 1; j < sz; j++ {
				o := points[ms[j][0]]
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
	xpo, ypo := p[0]-o[0], p[1]-o[1]
	xqo, yqo := q[0]-o[0], q[1]-o[1]
	if xpo*xqo+ypo*yqo != 0 { // not rectangle
		return math.MaxFloat64
	}
	return length(xpo, ypo) * length(xqo, yqo)
}

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
