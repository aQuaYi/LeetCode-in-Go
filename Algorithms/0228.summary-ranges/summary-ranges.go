package Problem0228

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
	for i := 1; i < l; i++ {
		if a[i-1]+1 != a[i] {
			if a[i-1] == begin {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, a[i-1])
			}
			begin = a[i]
			res = append(res, str)
		}
	}
	if a[l-1] == begin {
		str = fmt.Sprintf("%d", begin)
	} else {
		str = fmt.Sprintf("%d->%d", begin, a[l-1])
	}
	res = append(res, str)

	return res
}
