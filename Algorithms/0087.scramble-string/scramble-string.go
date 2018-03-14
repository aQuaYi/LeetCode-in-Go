package problem0087

func isScramble(a, b string) bool {
	// 相等的话，则为true
	if a == b {
		return true
	}

	n := len(a)

	// 检查 两个字符串是否具有相同的字符
	rec := make([]int, 256)
	for i := 0; i < n; i++ {
		rec[a[i]]++
		rec[b[i]]--
	}
	for i := range rec {
		if rec[i] != 0 {
			return false
		}
	}

	// 检查子字符串是否 scramble
	for i := 1; i < n; i++ {
		// 如果两个string中间某一个点，左边的substrings为scramble string，同时右边的substrings也为scramble string，则为true
		if isScramble(a[0:i], b[0:i]) &&
			isScramble(a[i:], b[i:]) {
			return true
		}

		// 如果两个string中间某一个点，s1左边的substring和s2右边的substring为scramble string, 同时s1右边substring和s2左边的substring也为scramble string，则为true
		if isScramble(a[0:i], b[n-i:]) &&
			isScramble(a[i:], b[0:n-i]) {
			return true
		}
	}

	return false
}
