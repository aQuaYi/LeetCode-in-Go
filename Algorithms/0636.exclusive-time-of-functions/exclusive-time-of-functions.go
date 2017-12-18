package Problem0636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	size := len(logs)
	stack := make([][]int, 0, size)
	res := make([]int, n)
	startTime := 0

	for _, log := range logs {
		soe, l := analyze(log)

		if soe == "start" {
			stack = append(stack, l)
			continue
		}

		if len(stack) > 0 {
			startTime = stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
		}

		res[l[0]] += l[1] - startTime + 1

		startTime = l[1]
	}

	return res
}

func analyze(log string) (string, []int) {
	ss := strings.Split(log, ":")
	funcID, _ := strconv.Atoi(ss[0])
	timeStamp, _ := strconv.Atoi(ss[2])
	return ss[1], []int{funcID, timeStamp}
}
