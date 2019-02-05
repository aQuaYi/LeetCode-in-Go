package problem0939

import "fmt"

const initialArea = 2000000000 // > 40000*40000

func minAreaRect(points [][]int) int {
	isExisted := make(map[string]bool, 500)
	xs := make(map[int][]int, 500)
	ys := make(map[int][]int, 500)
	for _, p := range points {
		px, py := p[0], p[1]
		isExisted[convert(px, py)] = true
		xs[px] = append(xs[px], py)
		ys[py] = append(ys[py], px)
	}

	area := initialArea

	for _, p := range points {
		xld, yld := p[0], p[1]
		for _, yrt := range xs[xld] {
			for _, xrt := range ys[yld] {
				if xld >= xrt || yld >= yrt {
					continue
				}
				if isExisted[convert(xrt, yrt)] {
					fmt.Printf("(%d,%d) (%d,%d)\n", xld, yld, xrt, yrt)
					area = min(area, (xrt-xld)*(yrt-yld))
				}
			}
		}

	}

	if area == initialArea {
		return 0
	}
	return area
}

func convert(px, py int) string {
	return fmt.Sprintf("%d-%d", px, py)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
