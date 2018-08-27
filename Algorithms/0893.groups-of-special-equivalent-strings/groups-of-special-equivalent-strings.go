package problem0893

import (
	"sort"
	"strings"
)

func numSpecialEquivGroups(A []string) int {
	gMap := make(map[string]bool, len(A))
	for _, a := range A {
		gMap[encode(a)] = true
	}
	return len(gMap)
}

func encode(str string) string {
	evens := make([]string, 0, 20)
	odds := make([]string, 0, 20)
	for i := range str {
		if i%2 == 0 {
			evens = append(evens, str[i:i+1])
		} else {
			odds = append(odds, str[i:i+1])
		}
	}

	sort.Strings(evens)
	sort.Strings(odds)

	return strings.Join(evens, "") + strings.Join(odds, "")
}
