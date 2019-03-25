package problem0963

import "math"

func minAreaFreeRect(points [][]int) float64 {
	n := len(points)
	res := math.MaxFloat64
	m := map[int]map[int]struct{}{}
	for _, p := range points {
		if mm, ok := m[p[0]]; ok {
			mm[p[1]] = struct{}{}
		} else {
			m[p[0]] = map[int]struct{}{p[1]: struct{}{}}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				p1, p2, p3 := points[i], points[j], points[k]
				x1, y1 := p2[0]-p1[0], p2[1]-p1[1]
				x2, y2 := p3[0]-p1[0], p3[1]-p1[1]
				if x1*x2+y1*y2 != 0 {
					continue
				}

				if _, ok := m[p3[0]-p1[0]+p2[0]][p3[1]-p1[1]+p2[1]]; !ok {
					continue
				}
				area := math.Sqrt(float64(x1*x1+y1*y1)) * math.Sqrt(float64(x2*x2+y2*y2))
				res = math.Min(res, area)
			}
		}
	}

	if res == math.MaxFloat64 {
		return 0
	}
	return res
}
