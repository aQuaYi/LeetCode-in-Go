package problem0870

import (
	"sort"
)

func advantageCount(A []int, B []int) []int {
	size := len(A)
	tmp := make([][2]int, size)

	for i, n := range B {
		tmp[i][0], tmp[i][1] = n, i
	}

	sort.Slice(tmp, func(i int, j int) bool {
		return tmp[i][0] < tmp[j][0]
	})

	res := make([]int, size)

	sort.Ints(A)

	l, r := 0, size-1

	for _, a := range A {
		if tmp[l][0] < a {
			res[tmp[l][1]] = a
			l++
		} else {
			res[tmp[r][1]] = a
			r--
		}
	}

	return res
}
