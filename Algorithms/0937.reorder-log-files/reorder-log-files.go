package problem0937

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	size := len(logs)

	letters := make([][]string, 0, size)
	digits := make([]string, 0, size)

	for _, log := range logs {
		ls := strings.SplitN(log, " ", 2)
		b := ls[1][0]
		if '0' <= b && b <= '9' {
			digits = append(digits, log)
		} else {
			letters = append(letters, ls)
		}
	}

	sort.Slice(letters, func(i int, j int) bool {
		return letters[i][1] < letters[j][1]
	})

	res := make([]string, 0, size)
	for _, ls := range letters {
		res = append(res, strings.Join(ls, " "))
	}

	res = append(res, digits...)

	return res
}
