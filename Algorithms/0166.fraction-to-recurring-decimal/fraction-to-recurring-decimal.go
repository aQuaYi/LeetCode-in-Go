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
		// n, d 异号，则结果前面有负号
		return "-" + fractionToDecimal(abs(n), abs(d))
	}

	// 确保 n  和 d 是非负数
	n, d = abs(n), abs(d)

	// n / d 的小数部分
	ds := ""
	if n >= d {
		ds = fractionToDecimal(n%d, d)
		return strconv.Itoa(n/d) + ds[1:]
	}

	// rec 记录所有出现过的 n
	rec := make(map[int]int, 1024)
	// idx 是 n/d 的结果的在 ds 的索引号
	idx := 0
	for {
		if i, ok := rec[n]; ok {
			// n 重复出现，则说明出现了循环的部分
			// 循环部分的起点，就是 n 上次出现的 idx 值
			return fmt.Sprintf("0.%s(%s)", ds[:i], ds[i:])
		}

		rec[n] = idx

		n *= 10
		idx++

		ds += string(n/d + '0')
		n %= d

		if n == 0 {
			// 没有循环部分
			return "0." + ds
		}
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
