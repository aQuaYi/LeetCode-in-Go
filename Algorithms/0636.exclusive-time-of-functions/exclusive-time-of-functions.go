package Problem0636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	size := len(logs)
	stack := make([][]int, 0, size)
	res := make([]int, n)
	time := 0
	for _, log := range logs {
		soe, l := analyze(log)

		if soe == "start" {
			stack = append(stack, l)
			continue
		}

		ls := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fnTime := l[1] - ls[1]

		res[ls[0]] += fnTime - time

		if len(stack) > 0 {
			time = fnTime
		} else {
			time = 0
		}
	}

	return res
}

func analyze(log string) (string, []int) {
	ss := strings.Split(log, ":")
	funcID, _ := strconv.Atoi(ss[0])
	timeStamp, _ := strconv.Atoi(ss[2])
	return ss[1], []int{funcID, timeStamp}
}
