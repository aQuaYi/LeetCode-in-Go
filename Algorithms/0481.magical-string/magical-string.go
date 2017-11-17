package Problem0481

import (
	"strings"
)

func magicalString(n int) int {
	bytes := make([]byte, n+5)
	copy(bytes, "122")

	i, j := 2, 2
	for {
		// b 是下一个要填写的字符
		b := (bytes[j] - '0') ^ 3 + '0'
		// c 是下一个要填写字符的数量
		c := bytes[i] - '0'
		i++

		for c > 0 {
			j++
			bytes[j] = b
			c--
		}

		if j >= n {
			break
		}
	}

	return strings.Count(string(bytes[:n]), "1")
}
