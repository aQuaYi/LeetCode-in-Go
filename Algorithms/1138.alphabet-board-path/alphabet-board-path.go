package problem1138

import "strings"

func alphabetBoardPath(target string) string {
	var sb strings.Builder

	// form [x1,y1] move to [x2,y2]
	move := func(x1, y1, x2, y2 int) {
		for y1 > y2 {
			y1--
			sb.WriteByte('L')
		}
		for x1 > x2 {
			x1--
			sb.WriteByte('U')
		}
		for x1 < x2 {
			x1++
			sb.WriteByte('D')
		}
		for y1 < y2 {
			y1++
			sb.WriteByte('R')
		}
		return
	}

	x1, y1 := 0, 0
	for _, r := range target {
		x2, y2 := coordinate(r)
		move(x1, y1, x2, y2)
		sb.WriteByte('!')
		x1, y1 = x2, y2
	}

	return sb.String()
}

func coordinate(r rune) (x, y int) {
	a := int(r - 'a')
	x, y = a/5, a%5
	return
}
