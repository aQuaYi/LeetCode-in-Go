package problem0821

func shortestToChar(S string, C byte) []int {
	res := make([]int, len(S))
	left, right := -len(S), next(-1, C, S)
	for i := range S {
		if i > right {
			left = right
			right = next(right, C, S)
		}

		res[i] = min(i-left, right-i)
	}

	return res
}

func next(right int, c byte, s string) int {
	idx := right + 1
	for idx < len(s) {
		if s[idx] == c {
			return idx
		}
		idx++
	}
	// 由于 s[right+1:] 中不再含有 c
	// 返回的 idx 需要达到 2*len(s)
	// 来确保后续的 min 操作的正确性
	return len(s) * 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
