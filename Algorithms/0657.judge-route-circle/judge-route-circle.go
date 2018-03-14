package problem0657

import "strings"

func judgeCircle(moves string) bool {
	up := strings.Count(moves, "U")
	down := strings.Count(moves, "D")
	left := strings.Count(moves, "L")
	right := strings.Count(moves, "R")

	return up == down && left == right
}
