package problem0682

import (
	"strconv"
)

func calPoints(ops []string) int {
	pointStack := make([]int, 0, len(ops))

	for i := range ops {
		switch ops[i] {
		case "+":
			r1 := pointStack[len(pointStack)-1]
			r2 := pointStack[len(pointStack)-2]
			pointStack = append(pointStack, r1+r2)
		case "D":
			r1 := pointStack[len(pointStack)-1]
			pointStack = append(pointStack, 2*r1)
		case "C":
			pointStack = pointStack[:len(pointStack)-1]
		default:
			point, _ := strconv.Atoi(ops[i])
			pointStack = append(pointStack, point)
		}
	}

	res := 0
	for _, p := range pointStack {
		res += p
	}
	return res
}
