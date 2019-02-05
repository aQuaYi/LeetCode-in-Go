package problem0939

import "fmt"

const initialArea = 2000000000 // > 40000*40000

func minAreaRect(points [][]int) int {
	isExisted := make(map[string]bool, 500)
	for _, p := range points {
		px, py := p[0], p[1]
		isExisted[convert(px, py)] = true
	}

	area := initialArea

	for _, pld := range points {
		xld, yld := pld[0], pld[1]
		for _, prt := range points {
			xrt, yrt := prt[0], prt[1]
			if xld >= xrt || yld >= yrt {
				continue
			}
			if isExisted[convert(xrt, yld)] && isExisted[convert(xld, yrt)] {
				area = min(area, (xrt-xld)*(yrt-yld))
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
