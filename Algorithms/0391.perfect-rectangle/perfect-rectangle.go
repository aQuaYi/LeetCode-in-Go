package Problem0391

import "fmt"

func isRectangleCover(rects [][]int) bool {
	if len(rects) == 0 || len(rects[0]) == 0 {
		return false
	}

	x1, x2 := 1<<63-1, -1<<63
	y1, y2 := x1, x2

	set := make(map[string]bool)
	area := 0

	var s1, s2, s3, s4 string

	for _, r := range rects {
		x1 = min(r[0], x1)
		y1 = min(r[1], y1)
		x2 = max(r[2], x2)
		y2 = max(r[3], y2)

		area += (r[2] - r[0]) * (r[3] - r[1])

		s1 = fmt.Sprintf("%d %d", r[0], r[1])
		s2 = fmt.Sprintf("%d %d", r[0], r[3])
		s3 = fmt.Sprintf("%d %d", r[2], r[3])
		s4 = fmt.Sprintf("%d %d", r[2], r[1])

		if set[s1] {
			delete(set, s1)
		} else {
			set[s1] = true
		}

		if set[s2] {
			delete(set, s2)
		} else {
			set[s2] = true
		}

		if set[s3] {
			delete(set, s3)
		} else {
			set[s3] = true
		}

		if set[s4] {
			delete(set, s4)
		} else {
			set[s4] = true
		}

	}

	s1 = fmt.Sprintf("%d %d", x1, y1)
	s2 = fmt.Sprintf("%d %d", x1, y2)
	s3 = fmt.Sprintf("%d %d", x2, y1)
	s4 = fmt.Sprintf("%d %d", x2, y2)

	if !set[s1] ||
		!set[s2] ||
		!set[s3] ||
		!set[s4] ||
		len(set) != 4 {
		return false
	}

	return area == (x2-x1)*(y2-y1)
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
