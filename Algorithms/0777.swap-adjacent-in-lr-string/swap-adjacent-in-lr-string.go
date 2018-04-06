package problem0777

import (
	"strings"
)

func canTransform(start string, end string) bool {
	return strings.Replace(start, "X", "", -1) == strings.Replace(end, "X", "", -1) &&
		isOK(idxs(start, 'R'), idxs(end, 'R'), isLessOrEqual) &&
		isOK(idxs(start, 'L'), idxs(end, 'L'), isMoreOrEqual)
}

// idxs 返回 s 中所有 b 字符的索引号
func idxs(s string, b byte) []int {
	res := make([]int, 0, len(s))
	for i := range s {
		if s[i] == b {
			res = append(res, i)
		}
	}
	return res
}

// 对于任意的 i ，都能使的 isAvailable(a[i], b[i]) 返回 true
// 则 isOK 返回 true
func isOK(a, b []int, isAvailable func(x, y int) bool) bool {
	for i := range a {
		if !isAvailable(a[i], b[i]) {
			return false
		}
	}
	return true
}

func isLessOrEqual(x, y int) bool {
	if x <= y {
		return true
	}
	return false
}

func isMoreOrEqual(x, y int) bool {
	if x >= y {
		return true
	}
	return false
}
