package problem0442

import (
	"sort"
)

func findDuplicates(a []int) []int {
	for i := 0; i < len(a); i++ {
		for a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	res := make([]int, 0, len(a))

	for i, n := range a {
		if i != n-1 {
			res = append(res, n)
		}
	}

	sort.Ints(res)

	return res
}
