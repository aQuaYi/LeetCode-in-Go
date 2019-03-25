package problem0963

import "math"

func minAreaFreeRect(points [][]int) float64 {
	size := len(points)
	res := math.MaxFloat64

	check := make(map[[2]int]struct{}, size)
	for _, p := range points {
		check[[2]int{p[0], p[1]}] = struct{}{}
	}

	for i := 0; i < size; i++ {
		o := points[i]
		for j := i + 1; j < size; j++ {
			p := points[j]
			x1, y1 := p[0]-o[0], p[1]-o[1] // vector form o to p
			for k := j + 1; k < size; k++ {
				q := points[k]
				x2, y2 := q[0]-o[0], q[1]-o[1] // vector form o to q

				if x1*x2+y1*y2 != 0 { // not a rectangle
					continue
				}

				p4 := [2]int{q[0] - o[0] + p[0], q[1] - o[1] + p[1]}
				if _, ok := check[p4]; !ok { // miss fourth point
					continue
				}

				area := length(x1, y1) * length(x2, y2)
				res = math.Min(res, area)
			}
		}
	}

	if res == math.MaxFloat64 {
		return 0
	}
	return res
}

// length of vector
func length(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}
