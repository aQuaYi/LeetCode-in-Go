package problem1138

import "strings"

func alphabetBoardPath(target string) string {
	x1, y1 := 0, 0
	var sb strings.Builder
	for _, r := range target {
		x2, y2 := coordinate(r)
		sb.WriteString(move(x1, y1, x2, y2))
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

// move form [x1,y1] to [x2,y2]
func move(x1, y1, x2, y2 int) string {
	if x1 == x2 && y1 == y2 {
		return ""
	}
	var sb strings.Builder
	if x1 == 5 {
		sb.WriteByte('U')
		x1--
	}
	for y1 < y2 {
		y1++
		sb.WriteByte('R')
	}
	for y1 > y2 {
		y1--
		sb.WriteByte('L')
	}
	for x1 < x2 {
		x1++
		sb.WriteByte('D')
	}
	for x1 > x2 {
		x1--
		sb.WriteByte('U')
	}
	return sb.String()
}
