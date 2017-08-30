package Problem0611

import (
	"sort"
)

func triangleNumber(a []int) int {
	sort.Ints(a)
	n := len(a)

	res := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j] > a[k] {
					res++
				}
			}
		}
	}

	return res
}
