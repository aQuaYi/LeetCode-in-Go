package Problem0402

import (
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	if k == 0 {
		i := 0
		for i < len(num)-1 && num[i] == '0' {
			i++
		}
		return num[i:]
	}

	var i int
	for i+1 < len(num) && num[i] <= num[i+1] {
		i++
	}
	return removeKdigits(strings.Replace(num, num[i:i+1], "", 1), k-1)
}
