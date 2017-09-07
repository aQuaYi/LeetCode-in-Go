package Problem0087

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	// 题目规定了， s1 与 s2 等长
	n := len(s1)
	if n == 1 {
		return false
	}

	dp := make ([][][]bool, n)


	return true
}
