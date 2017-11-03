package Problem0402

import (
	"strings"
)

func minNum(s string) (at int) {
	for i := 1; i < len(s); i++ {
		if s[i] < s[at] {
			at = i
		}
	}
	return
}

func removeKdigits(num string, k int) string {
	var prefix string
	for k > 0 {
		if len(num) <= k {
			num = ""
			break
		}
		pos := strings.IndexByte(num, '0')
		if pos > 0 && pos <= k {
			k -= pos
			num = num[pos:]
			num1 := strings.TrimLeft(num, "0")
			if prefix != "" {
				prefix += num[:len(num)-len(num1)]
			}
			num = num1
			continue
		}

		pos = minNum(num[:k+1])
		k -= pos
		prefix += string(num[pos])
		num = num[pos+1:]

	}
	num = prefix + num
	if len(num) == 0 {
		return "0"
	}
	return num
}
