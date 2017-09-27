package Problem0166

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(n int, d int) string {
	if n == 0 {
		return "0"
	}

	if n*d < 0 {
		return "-" + fractionToDecimal(abs(n), abs(d))
	}

	// 确保 n  和 d 是非负数
	n, d = abs(n), abs(d)

	// n / d 的小数部分
	ds := ""

	if n < d {
		rec := make(map[int]int, 1024)
		idx := 0
		for {
			if i, ok := rec[n]; ok {
				return fmt.Sprintf("0.%s(%s)", ds[:i], ds[i:])
			}

			rec[n] = idx

			n *= 10
			idx++

			ds += string(n/d + '0')
			n %= d

			if n == 0 {
				return "0." + ds
			}
		}
	}

	ds = fractionToDecimal(n%d, d)
	return strconv.Itoa(n/d) + ds[1:]
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
