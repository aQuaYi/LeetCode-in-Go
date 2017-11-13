package Problem0455

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(g)))
	sort.Ints(s)

	n := len(s)

	content := 0
	for _, greed := range g {
		if len(s) == 0 {
			break
		}
		if greed <= s[n-1] {
			n--
			content++
		}
	}
	return content
}
