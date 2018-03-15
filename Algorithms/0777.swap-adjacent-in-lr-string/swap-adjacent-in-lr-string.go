package problem0777

func canTransform(start string, end string) bool {
	return isOK(indexSlice(start, 'R'), indexSlice(end, 'R'), notBigger) &&
		isOK(indexSlice(start, 'L'), indexSlice(end, 'L'), notSmaller)
}

// indexSlice 返回 s 中所有 b 字符的索引号
func indexSlice(s string, b byte) []int {
	res := make([]int, 0, len(s))
	for i := range s {
		if s[i] == b {
			res = append(res, i)
		}
	}
	return res
}

func isOK(a, b []int, compare func(x, y int) bool) bool {
	for i := range a {
		if compare(a[i], b[i]) {
			return false
		}
	}
	return true
}

func notBigger(x, y int) bool {
	if x <= y {
		return true
	}
	return false
}

func notSmaller(x, y int) bool {
	if x >= y {
		return true
	}
	return false
}
