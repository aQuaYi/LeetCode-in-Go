package problem0937

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s1 := strings.SplitN(logs[i], " ", 2)
		s2 := strings.SplitN(logs[j], " ", 2)
		f1, f2 := "0"+s1[1], "0"+s2[1]
		if unicode.IsNumber(rune(f1[1])) {
			f1 = "1"
		}
		if unicode.IsNumber(rune(f2[1])) {
			f2 = "1"
		}
		return f1 < f2
	})
	return logs
}
