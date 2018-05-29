package problem0828

func uniqueLetterString(s string) int {
	size := len(s)
	res := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j <= size; j++ {
			res += uniqueNum(s[i:j])
		}
	}
	return res
}

func uniqueNum(s string) int {
	isSaw := map[rune]bool{}
	res := 0
	for _, b := range s {
		if isSaw[b] {
			res--
		} else {
			isSaw[b] = true
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
