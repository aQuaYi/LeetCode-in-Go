package problem0937

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	size := len(logs)

	letters := make([]string, 0, size)
	digits := make([]string, 0, size)

	for _, log := range logs {
		if isDigit(log) {
			digits = append(digits, log)
		} else {
			letters = append(letters, log)
		}
	}

	sort.Slice(letters, func(i int, j int) bool {
		li, lj := letters[i], letters[j]
		li = li[strings.Index(li, " "):]
		lj = lj[strings.Index(lj, " "):]
		return strings.Compare(li, lj) < 0
	})

	return append(letters, digits...)
}

func isDigit(log string) bool {
	b := log[strings.Index(log, " ")+1]
	return '0' <= b && b <= '9'
}
