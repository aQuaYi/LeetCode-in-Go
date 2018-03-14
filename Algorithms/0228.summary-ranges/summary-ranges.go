package problem0228

import (
	"fmt"
)

func summaryRanges(a []int) []string {
	res := []string{}

	l := len(a)
	if l == 0 {
		return res
	}

	begin := a[0]
	str := ""

	for i := 0; i < l; i++ {
		if i == l-1 || a[i]+1 != a[i+1] {
			if a[i] == begin {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, a[i])
			}

			if i+1 < l {
				begin = a[i+1]
			}

			res = append(res, str)
		}
	}

	return res
}
