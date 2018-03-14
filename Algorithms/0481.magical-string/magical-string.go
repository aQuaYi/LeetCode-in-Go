package problem0481

import (
	"strings"
)

// 解题思路
// 1. 构造出完美字符串 s
// 2. 统计 s 中 1 的个数
func magicalString(n int) int {
	// n+2 是为了防止添加字符时溢出
	bytes := make([]byte, n+2)

	copy(bytes, "122")

	// i 指向下一个要添加字符的数量
	// j 始终指向 bytes 中最后一个字符，以便确定下一个需要添加的字符
	i, j := 2, 2
	for j < n {
		// b 是下一个要填写的字符
		b := (bytes[j] - '0') ^ 3 + '0'
		// c 是下一个要填写字符的数量
		c := bytes[i] - '0'
		i++

		// 在bytes中添加 c 个字符 b
		for c > 0 {
			j++
			bytes[j] = b
			c--
		}
	}

	return strings.Count(string(bytes[:n]), "1")
}
