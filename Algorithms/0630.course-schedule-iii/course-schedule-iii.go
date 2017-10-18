package Problem0630

import (
	"sort"
)

func scheduleCourse(courses [][]int) int {
	css := cs(courses)
	sort.Sort(css)

	var count, now int
	for _, c := range css {
		if now+c[0] <= c[1] {
			now += c[0]
			count++
		}
	}

	return count
}

type cs [][]int

func (c cs) Len() int {
	return len(c)
}

func (c cs) Less(i, j int) bool {
	return c[i][1] < c[j][1]
}

func (c cs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
