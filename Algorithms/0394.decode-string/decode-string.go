package problem0394

import (
	"strconv"
)

func decodeString(s string) string {
	n := len(s)

	// i 是第一个数字的位置
	i := 0
	for i < n && (s[i] < '0' || '9' < s[i]) {
		i++
	}
	if i == n {
		// 没有数字，直接返回 s
		return s
	}

	// j 是第一个 '[' 的位置
	j := i + 1
	// 由题意可知，s 很规范
	// 存在数字的话，必定存在 '[' 和 ']'
	for s[j] != '[' {
		j++
	}

	// k 是与 j 的 '[' 对应的 ']' 的位置
	k := j
	count := 1
	for count > 0 {
		k++

		if s[k] == '[' {
			count++
		} else if s[k] == ']' {
			count--
		}
	}

	//      i：第一个数字的位置
	//      |  j：第一个 '[' 的位置
	//      |  |       k：与 j 的 '[' 对应的 ']' 的位置
	//      ↓  ↓       ↓
	// "abcd234[*******]efg"

	// 题目说了， s 很规范
	num, _ := strconv.Atoi(s[i:j])

	return s[:i] + times(num, decodeString(s[j+1:k])) + decodeString(s[k+1:])
}

func times(n int, s string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
